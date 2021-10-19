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
	application := framework.NewApp(os.DirFS(path.BasePath("config")), path.BasePath("storage/logs"))
	app.BootstrapApp(application)

	e := application.Run()
	if e != nil {
		log.Fatal(e)
	}
}
