package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/enorith/enorith/config"
	"github.com/enorith/enorith/internal/app"
	"github.com/enorith/enorith/internal/app/routes"
	"github.com/enorith/framework"
	"github.com/enorith/http"
	"github.com/enorith/http/router"
	"github.com/joho/godotenv"
)

const ServeAt = ":3113"

func main() {
	application := framework.NewApp(config.FS)
	app.BootstrapApp(application)


	if application.GetConfig().Env != "production" {
		cwd, _ := os.Getwd()
		env := filepath.Join(cwd, ".env")
		log.Printf("loading .env [%s]", env)
		_ = godotenv.Load(env)		
	}

	e := application.Run(ServeAt, func(rw *router.Wrapper, k *http.Kernel) {
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
