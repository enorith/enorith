package api

import (
	"github.com/enorith/authenticate"
	"github.com/enorith/authenticate/jwt"
	"github.com/enorith/enorith/internal/app/requests"
	"github.com/enorith/enorith/internal/pkg/auth"
)

type AuthHandler struct {
}

func (AuthHandler) Login(r requests.LoginRequest, g authenticate.Guard, p auth.UserProvider) (jwt.Token, error) {

	u, e := p.Attempt(r)
	if e != nil {
		return jwt.Token{}, e
	}
	e = g.Auth(u)

	return g.(*jwt.Guard).Token(), e
}

func (AuthHandler) User(g authenticate.Guard) authenticate.User {
	return g.User()
}
