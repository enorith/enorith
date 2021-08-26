package auth

import (
	"github.com/enorith/container"
	"github.com/enorith/framework"
	"github.com/enorith/http/contracts"
)

type Service struct {
}

//Register service when app starting, before http server start
// you can configure service, prepare global vars etc.
// running at main goroutine
func (s Service) Register(app *framework.App) error {
	// tx, e := gormdb.DefaultManager.GetConnection()
	// if e != nil {
	// 	return e
	// }
	// provider := UserProvider{tx}

	return nil
}

//Lifetime container callback
// usually register request lifetime instance to IoC-Container (per-request unique)
// this function will run before every request handling
func (s Service) Lifetime(ioc container.Interface, request contracts.RequestContract) {
	panic("not implemented") // TODO: Implement
}
