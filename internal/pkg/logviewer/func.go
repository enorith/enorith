package logviewer

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/enorith/supports/collection"
	"github.com/icza/backscanner"
	jsoniter "github.com/json-iterator/go"
)

var groups = make([]LogGroup, 0)

func RegisterGroup(g LogGroup) {
	groups = append(groups, g)
}

func GetGroupByDir(dir string) LogGroup {
	for _, g := range groups {
		if g.Dir() == dir {
			return g
		}
	}
	return NewSimpleDirGroup(dir, dir)
}

type ListGroupItem struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
}

func ListGroups() []ListGroupItem {
	return collection.Map(collection.SortBy(groups, func(a, b LogGroup) bool {
		return a.Order() > b.Order()
	}), func(g LogGroup) ListGroupItem {
		return ListGroupItem{Name: g.GroupName(), Dir: g.Dir()}
	})
}

type ListGroupFileItem struct {
	Title    string `json:"title"`
	Filename string `json:"filename"`
}

func ListGroupFilesFromDir(dir string) []ListGroupFileItem {
	lg := GetGroupByDir(dir)
	var items []ListGroupFileItem
	if lg != nil {
		filepath.WalkDir(filepath.Join(LogDir, dir), func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				items = append(items, ListGroupFileItem{Title: lg.ResolveTitle(strings.ReplaceAll(d.Name(), ".log", "")), Filename: d.Name()})
			}
			return nil
		})
	}

	return collection.SortBy(items, func(a, b ListGroupFileItem) bool {
		return a.Title > b.Title
	})
}

type LogPageMeta struct {
	Pos  int64 `json:"pos"`
	Line int   `json:"line"`
	End  bool  `json:"end"`
}

type LogList struct {
	Data       []map[string]interface{} `json:"data"`
	ContentKey string                   `json:"content"`
	Search     []string                 `json:"search"`
	Meta       LogPageMeta              `json:"meta"`
}

func GetLogs(dir, filename string, pos int64) LogList {
	maxLine := 50
	var logs []map[string]interface{}

	file := filepath.Join(LogDir, dir, filename)

	f, e := os.OpenFile(file, os.O_RDONLY, 0)
	lg := GetGroupByDir(dir)
	var parser LogParser
	contentKey := "msg"
	searchKeys := []string{"msg"}
	if lg != nil {
		name := strings.Split(filename, ".")[0]
		var ok bool
		parser, ok = lg.Parsers()[name]
		if !ok {
			parser = CommonLogParser{}
		}
		if ck, is := parser.(WithContentKey); is {
			contentKey = ck.ContentKey()
		}

		if sk, is := parser.(WithSearchKeys); is {
			searchKeys = sk.SearchKeys()
		}

	}
	meta := LogPageMeta{}
	meta.Pos = pos
	line := 0
	if e == nil {

		fi, _ := f.Stat()

		defer f.Close()
		//f.Seek(pos, io.SeekEnd)
		//scanner := bufio.NewScanner(f)
		scanner := backscanner.New(f, int(fi.Size()-pos))
		for {

			bs, _, err := scanner.LineBytes()
			if err != nil {
				// if err == io.EOF {
				// 	fmt.Printf("%q is not found in file.\n", what)
				// } else {
				// 	fmt.Println("Error:", err)
				// }
				break
			}

			if line >= maxLine {
				break
			}

			line++
			var log map[string]interface{}
			//bs := scanner.Bytes()
			meta.Pos += int64(len(bs))

			e = jsoniter.Unmarshal(bs, &log)
			if e == nil {
				lm := parser.LabelMap()
				item := make(map[string]interface{})
				for k, v := range log {
					l, y := lm[k]
					if !y {
						l = k
					}
					item[l] = v
				}

				logs = append(logs, parser.ResolveLog(item))
			}
		}
	}
	//utils.Reverse(logs)
	meta.Line = line
	meta.End = line == 0
	return LogList{
		Data:       logs,
		ContentKey: contentKey,
		Search:     searchKeys,
		Meta:       meta,
	}
}
