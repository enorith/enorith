package api

import (
	"github.com/enorith/authenticate"
	"github.com/enorith/authenticate/jwt"
	"github.com/enorith/enorith/internal/app/models"
	"github.com/enorith/enorith/internal/app/requests"
	"github.com/enorith/enorith/internal/pkg/auth"
	"github.com/enorith/framework/authentication"
	"gorm.io/gorm"
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

func (AuthHandler) Register(user models.User, tx *gorm.DB, g authenticate.Guard) (t jwt.Token, e error) {
	user.Password, e = authentication.Hash(user.Password)

	if e != nil {
		return jwt.Token{}, e
	}
	tx.Save(&user)
	g.Auth(user)
	t = g.(*jwt.Guard).Token()

	return
}
