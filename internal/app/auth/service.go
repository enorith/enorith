package auth

import (
	"github.com/enorith/container"
	"github.com/enorith/framework"
	"github.com/enorith/framework/authentication"
	"github.com/enorith/gormdb"
	"github.com/enorith/http/contracts"
	"github.com/enorith/http/errors"
	"github.com/enorith/language"
)

type Service struct {
}

//Register service when app starting, before http server start
// you can configure service, prepare global vars etc.
// running at main goroutine
func (s Service) Register(app *framework.App) error {
	tx, e := gormdb.DefaultManager.GetConnection()
	if e != nil {
		return e
	}
	msg, _ := language.T("auth", "failed")
	AuthFailedError = errors.UnprocessableEntity(msg)

	provider := NewUserProvider(tx)

	authentication.AuthManager.WithProvider("users", provider)

	return nil
}

//Lifetime container callback
// usually register request lifetime instance to IoC-Container (per-request unique)
// this function will run before every request handling
func (s Service) Lifetime(ioc container.Interface, request contracts.RequestContract) {
	ioc.BindFunc("middleware.auth", func(c container.Interface) (interface{}, error) {
		return c.Instance(Middleware{})
	}, true)
}
