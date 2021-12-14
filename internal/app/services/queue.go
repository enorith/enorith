package services

import (
	"github.com/enorith/framework"
)

type QueueService struct {
}

//Register service when app starting, before http server start
// you can configure service, prepare global vars etc.
// running at main goroutine
func (qs QueueService) Register(app *framework.App) error {
	// register your job handlers
	// queue.RegisterHandler(job.FooHandler)

	return nil
}
