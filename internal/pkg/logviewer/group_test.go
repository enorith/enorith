package logviewer_test

import (
	"github.com/enorith/enorith/internal/pkg/logviewer"
	"github.com/enorith/enorith/internal/pkg/test"
	"github.com/enorith/enorith/internal/pkg/utils"
	"testing"
)

func TestGroup(t *testing.T) {
	test.Start()
	logviewer.RegisterGroup(logviewer.NewDatabaseGroup("database"))

	// utils.PrintJson(logviewer.ListGroups())
	utils.PrintJson(logviewer.ListGroupFilesFromDir("http"))
	// utils.PrintJson(logviewer.GetLogs("database", "database.2024-03-24.log"))
	// utils.PrintJson(logviewer.GetLogs("http", "http.2024-03-24.log"))
}
