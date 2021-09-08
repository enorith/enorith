package services

import (
	"github.com/enorith/container"
	"github.com/enorith/enorith/internal/app/routes"
	"github.com/enorith/framework"
	"github.com/enorith/framework/http/middleware"
	"github.com/enorith/http/contracts"
	"github.com/enorith/http/router"
)

type HttpService struct {
}

//Register service when app starting, before http server start
// you can configure service, prepare global vars etc.
// running at main goroutine
func (hs HttpService) Register(app *framework.App) error {

	return nil
}

func (hs HttpService) RegisterRoutes(rw *router.Wrapper) {
	// register web routes
	routes.WebRoutes(rw)

	// register api routes
	rw.Group(func(r *router.Wrapper) {
		routes.ApiRoutes(r)
	}, "api").Middleware("throttle.api")
}

//Lifetime container callback
// usually register request lifetime instance to IoC-Container (per-request unique)
// this function will run before every request handling
func (hs HttpService) Lifetime(ioc container.Interface, request contracts.RequestContract) {
	ioc.BindFunc("middleware.throttle.api", func(c container.Interface) (interface{}, error) {
		return middleware.Throttle(1, 60), nil
	}, false)
}
