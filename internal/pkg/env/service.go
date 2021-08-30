package env

import (
	"log"
	"os"
	"path"

	"github.com/enorith/container"
	"github.com/enorith/framework"
	"github.com/enorith/http/contracts"
	"github.com/joho/godotenv"
)

type Service struct {
}

//Register service when app starting, before http server start
// you can configure service, prepare global vars etc.
// running at main goroutine
func (s Service) Register(app *framework.App) error {
	env := app.GetEnv()

	cwd, _ := os.Getwd()

	var files []string
	files = append(files, path.Join(cwd, ".env."+env))
	files = append(files, path.Join(cwd, ".env"))
	log.Printf("loading dotenv files %v", files)
	_ = godotenv.Load(files...)

	return nil
}

//Lifetime container callback
// usually register request lifetime instance to IoC-Container (per-request unique)
// this function will run before every request handling
func (s Service) Lifetime(ioc container.Interface, request contracts.RequestContract) {
}
