package main

import (
	"log"

	"github.com/enorith/enorith/config"
	"github.com/enorith/enorith/internal/app"
	"github.com/enorith/framework"
)

const ServeAt = ":3113"

func main() {
	application := framework.NewApp(config.FS)
	app.BootstrapApp(application)

	e := application.Run()
	if e != nil {
		log.Fatal(e)
	}
}
