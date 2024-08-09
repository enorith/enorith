package logviewer

import "github.com/enorith/framework"

var LogDir string

func WithLogViewer(app *framework.App, logDir string) {
	LogDir = logDir

	RegisterGroup(NewDatabaseGroup("database"))
	RegisterGroup(NewHttpGroup("http"))
}
