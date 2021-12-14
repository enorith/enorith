# cron tasks handlers dirctroy


## getting started
```go
package schedule

func TaskHandlerFoo(tx *gorm.DB) func() {} {
    // do something awsome
}

// in internal/app/services/schedule.go
//......
func (s ScheduleService) Register(app *framework.App) error {

	app.Resolving(func(s *gocron.Scheduler, tx *gorm.DB /* inject your pendencies*/) {
		// run tasks
		s.Every(5).Hours().Tag("foo").Do(schedule.TaskHandlerFoo(tx))
	})

	return nil
}


```