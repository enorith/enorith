package routes

import (
	"github.com/enorith/enorith/internal/app/handlers/api"
	"github.com/enorith/http/router"
)

func ApiRoutes(r *router.Wrapper) {
	var apiHandler api.AuthHandler

	r.Post("login", apiHandler.Login)

	r.Group(func(r *router.Wrapper) {
		r.Get("user", apiHandler.User)
	}).Middleware("auth")

}
