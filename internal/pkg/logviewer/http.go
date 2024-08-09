package logviewer

import "strings"

type HttpParser struct {
	CommonLogParser
}

func (p HttpParser) LabelMap() map[string]string {
	m := p.CommonLogParser.LabelMap()

	return m
}

func (p HttpParser) ContentKey() string {
	return "<span class='bg-green-500 text-white px-1 rounded-md'>:method</span> <span class='bg-slate-200 px-1 rounded-md'>:status_code</span> :path"
}

func (p HttpParser) ResolveLog(log map[string]interface{}) map[string]interface{} {

	_, ok := log["method"]
	if !ok {
		log["method"] = "GET"
	}

	return log
}

func (p HttpParser) SearchKeys() []string {
	return []string{"path"}
}

type HttpGroup struct {
	WithDirGroup
}

func (dg *HttpGroup) GroupName() string {
	return "服务器请求日志"
}

func (dg *HttpGroup) ResolveTitle(filename string) string {
	filename = strings.ReplaceAll(filename, "http", "服务器请求")
	return filename
}

func (dg *HttpGroup) Parsers() map[string]LogParser {
	return map[string]LogParser{
		"http": HttpParser{CommonLogParser: CommonLogParser{}},
	}
}

func (dg *HttpGroup) Order() int {
	return 0
}

func NewHttpGroup(dir string) *HttpGroup {
	return &HttpGroup{WithDirGroup(dir)}
}
