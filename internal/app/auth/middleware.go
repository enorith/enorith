package auth

import (
	"reflect"

	"github.com/enorith/authenticate"
	"github.com/enorith/container"
	"github.com/enorith/framework/authentication"
	"github.com/enorith/http"
	"github.com/enorith/http/content"
	"github.com/enorith/http/contracts"
	"github.com/enorith/http/errors"
	"github.com/enorith/language"
)

type Middleware struct {
	authenticate.Guard
}

func (m Middleware) Handle(r contracts.RequestContract, next http.PipeHandler) contracts.ResponseContract {

	_, e := m.Check()
	if e != nil {
		msg, _ := language.T("auth", "unauthorized")
		return content.ErrResponseFromError(errors.Unauthorized(msg), 401, nil)
	}
	r.GetContainer().BindFunc(authentication.GuardType, func(c container.Interface) (reflect.Value, error) {
		return reflect.ValueOf(m.Guard), nil
	}, true)
	return next(r)
}
