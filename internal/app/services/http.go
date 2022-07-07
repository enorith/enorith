package services

import (
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/enorith/container"
	"github.com/enorith/enorith/internal/app/routes"
	"github.com/enorith/framework"
	"github.com/enorith/framework/http/middleware"
	"github.com/enorith/http/pipeline"
	"github.com/enorith/http/router"
)

type HttpService struct {
}

//Register service when app starting, before http server start
// you can configure service, prepare global vars etc.
// running at main goroutine
func (hs HttpService) Register(app *framework.App) error {

	app.Bind(func(ioc container.Interface) {
		ioc.BindFunc("middleware.api", func(c container.Interface) (interface{}, error) {
			return pipeline.MiddlewareChain(
				middleware.Throttle(1, 60),
			), nil
		}, false)
	})

	return nil
}

func (hs HttpService) RegisterRoutes(rw *router.Wrapper) {
	// register web routes
	rw.Group(func(r *router.Wrapper) {
		routes.WebRoutes(r)
	}).Middleware("session")

	// register api routes
	rw.Group(func(r *router.Wrapper) {
		routes.ApiRoutes(r)
	}, "api").Middleware("api")

	printRoutes(rw)
}

func printRoutes(rw *router.Wrapper) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Method"},
			{Align: simpletable.AlignCenter, Text: "Path"},
			{Align: simpletable.AlignCenter, Text: "Middleware"},
		},
	}
	for k, v := range rw.Routes() {
		for _, pr := range v {
			table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: k},
				{Align: simpletable.AlignLeft, Text: pr.Path()},
				{Align: simpletable.AlignLeft, Text: strings.Join(pr.Middleware(), ",")},
			})
		}
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
