package test

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/enorith/enorith/config"
	"github.com/enorith/enorith/internal/app"

	"github.com/enorith/framework"
	"github.com/enorith/supports/str"
	"github.com/joho/godotenv"
)

func Start() *framework.App {
	file := BasePath(".env")
	log.Printf("loading dotenv file %s", file)
	_ = godotenv.Load(file)
	os.Setenv("APP_ENV", "test")
	application := framework.NewApp(config.FS, LogDir())

	app.BootstrapApp(application, LogDir())

	application.Bootstrap()

	return application
}

func LogDir() string {
	return BasePath("storage/logs")
}

func BasePath(path ...string) string {
	base, _ := os.Getwd()

	if str.Contains(base, "\\internal") {
		base = strings.Split(base, "\\internal")[0]
	}

	if str.Contains(base, "/internal") {
		base = strings.Split(base, "/internal")[0]
	}

	paths := append([]string{base}, path...)

	return filepath.Join(paths...)
}
