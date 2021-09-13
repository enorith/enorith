package main

import (
	"log"
	"os"

	"github.com/enorith/enorith/internal/app"
	"github.com/enorith/enorith/internal/pkg/path"
	"github.com/enorith/framework"
)

const ServeAt = ":3113"

func main() {
	application := framework.NewApp(os.DirFS(path.BasePath("config")))
	app.BootstrapApp(application)

	e := application.Run()
	if e != nil {
		log.Fatal(e)
	}
}
