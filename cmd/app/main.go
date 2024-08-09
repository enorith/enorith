package main

import (
	"log"
	"os"

	"github.com/enorith/enorith/internal/app"
	"github.com/enorith/enorith/internal/pkg/env"
	"github.com/enorith/enorith/internal/pkg/path"
	"github.com/enorith/framework"
)

func main() {
	// load .env, before app created
	env.LoadDotenv()
	logDir := path.BasePath("storage/logs")
	application := framework.NewApp(os.DirFS(path.BasePath("config")), logDir)
	app.BootstrapApp(application, logDir)

	e := application.Run()
	if e != nil {
		log.Fatal(e)
	}
}
