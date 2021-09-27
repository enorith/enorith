package auth

import (
	"github.com/enorith/authenticate"
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

	_, e := m.Check()
	if e != nil {
		msg, _ := language.T("auth", "unauthorized")
		return content.ErrResponseFromError(errors.Unauthorized(msg), 401, nil)
	}
	return next(r)
}
