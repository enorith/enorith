package api

import (
	"github.com/enorith/authenticate"
	"github.com/enorith/authenticate/jwt"
	"github.com/enorith/enorith/internal/app/auth"
	"github.com/enorith/http/contracts"
)

type AuthHandler struct {
}

func (AuthHandler) Login(r contracts.RequestContract, g authenticate.Guard, p auth.UserProvider) (jwt.Token, error) {

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
