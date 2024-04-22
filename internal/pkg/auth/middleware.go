package auth

import (
	"github.com/enorith/authenticate"
	"github.com/enorith/container"
	"github.com/enorith/enorith/internal/app/models"
	"github.com/enorith/http/content"
	"github.com/enorith/http/contracts"
	"github.com/enorith/http/errors"
	"github.com/enorith/http/pipeline"
	"github.com/enorith/language"
)

type Middleware struct {
	authenticate.Guard
}

func (m Middleware) Handle(r contracts.RequestContract, next pipeline.PipeHandler) contracts.ResponseContract {

	user, e := m.Check()
	if e != nil {
		msg, _ := language.T("auth", "unauthorized")
		return content.ErrResponseFromError(errors.Unauthorized(msg), 401, nil)
	}

	if u, ok := user.(models.User); ok {
		r.GetContainer().BindFunc(AuthUser{}, func(c container.Interface) (interface{}, error) {
			return AuthUser{User: u}, nil
		}, true)
	}

	return next(r)
}
