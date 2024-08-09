package logviewer

type LogParser interface {
	LabelMap() map[string]string
	ResolveLog(map[string]interface{}) map[string]interface{}
}

type WithContentKey interface {
	ContentKey() string
}

type WithSearchKeys interface {
	SearchKeys() []string
}

type LogGroup interface {
	GroupName() string
	ResolveTitle(filename string) string
	Dir() string
	Parsers() map[string]LogParser
	Order() int
}

type CommonLogParser struct {
}

func (p CommonLogParser) LabelMap() map[string]string {
	return map[string]string{}
}

func (p CommonLogParser) ResolveLog(log map[string]interface{}) map[string]interface{} {
	return log
}

type WithDirGroup string

func (s WithDirGroup) Dir() string {
	return string(s)
}

type SimpleDirGroup struct {
	WithDirGroup
	name string
}

func (sd *SimpleDirGroup) GroupName() string {
	return sd.name
}

func (sd *SimpleDirGroup) ResolveTitle(filename string) string {
	return filename
}

func (sd *SimpleDirGroup) Parsers() map[string]LogParser {
	return map[string]LogParser{}
}

func (sd *SimpleDirGroup) Order() int {
	return 0
}

func NewSimpleDirGroup(dir, name string) *SimpleDirGroup {
	return &SimpleDirGroup{WithDirGroup(dir), name}
}
