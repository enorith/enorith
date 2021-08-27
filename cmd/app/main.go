package main

import (
	"fmt"
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
		k.OutputLog = true
		routes.WebRoutes(rw)
		rw.Group(func(r *router.Wrapper) {
			routes.ApiRoutes(r)
		}, "api")

		for k2, v := range rw.Routes() {
			for i, pr := range v {
				fmt.Println(k2, i, pr.Middleware(), pr.Path())
			}
		}
	})

	if e != nil {
		log.Fatal(e)
	}
}
