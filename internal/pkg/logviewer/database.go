package logviewer

import (
	"strings"

	"github.com/enorith/enorith/internal/pkg/utils"
)

type DatabaseParser struct {
	CommonLogParser
}

func (p DatabaseParser) LabelMap() map[string]string {
	m := p.CommonLogParser.LabelMap()
	m["duration"] = "duration"
	m["sql"] = "sql"
	m["sql_limit"] = "sql_limit"
	m["rows_affected"] = "影响行数"

	return m
}

func (p DatabaseParser) ContentKey() string {
	return "<span class='bg-green-500 text-white px-1 rounded-md'>:duration</span> :sql_limit"
}

func (p DatabaseParser) ResolveLog(log map[string]interface{}) map[string]interface{} {

	sql, ok := log["sql"].(string)
	if ok {
		log["sql_limit"] = utils.LimitString(sql, 120)
	}

	return log
}

func (p DatabaseParser) SearchKeys() []string {
	return []string{"sql"}
}

type DatabaseGroup struct {
	WithDirGroup
}

func (dg *DatabaseGroup) GroupName() string {
	return "数据库"
}

func (dg *DatabaseGroup) ResolveTitle(filename string) string {
	filename = strings.ReplaceAll(filename, "database", "数据库")
	return filename
}

func (dg *DatabaseGroup) Parsers() map[string]LogParser {
	return map[string]LogParser{
		"database": DatabaseParser{CommonLogParser: CommonLogParser{}},
	}
}

func (dg *DatabaseGroup) Order() int {
	return 0
}

func NewDatabaseGroup(dir string) *DatabaseGroup {
	return &DatabaseGroup{WithDirGroup(dir)}
}
