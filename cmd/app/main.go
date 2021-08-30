package main

import (
	"log"

	"github.com/enorith/enorith/config"
	"github.com/enorith/enorith/internal/app"
	"github.com/enorith/enorith/internal/app/routes"
	"github.com/enorith/framework"
	"github.com/enorith/http"
	"github.com/enorith/http/router"
)

const ServeAt = ":3113"

func main() {
	application := framework.NewApp(config.FS)
	app.BootstrapApp(application)

	e := application.Run(ServeAt, func(rw *router.Wrapper, k *http.Kernel) {
		k.KeepAlive()
		k.OutputLog = true
		routes.WebRoutes(rw)
		rw.Group(func(r *router.Wrapper) {
			routes.ApiRoutes(r)
		}, "api")
	})
	if e != nil {
		log.Fatal(e)
	}
}
