package routes

import (
	"github.com/enorith/enorith/internal/app/handlers/web"
	"github.com/enorith/http/router"
)

func WebRoutes(rw *router.Wrapper) {
	rw.Get("/", web.Index)
}
