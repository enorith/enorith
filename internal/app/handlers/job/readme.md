# queue job handlers dirctroy

## getting started
```go

package job

func FooHandler(foo payload.Foo, tx *gorm.DB //inject your dependencies) {
    // do something awsome
}

// in internal/app/services/queue.go
//......
func (qs QueueService) Register(app *framework.App) error {
	// register your job handlers
	queue.RegisterHandler(job.FooHandler)
    //......
	return nil
}

```