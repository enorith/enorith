package schedule

import (
	"github.com/enorith/framework"
	"github.com/go-co-op/gocron"
)

type Service struct {
}

//Register service when app starting, before http server start
// you can configure service, prepare global vars etc.
// running at main goroutine
func (s Service) Register(app *framework.App) error {
	app.Resolving(func(s *gocron.Scheduler) {
		// run tasks
		// s.Every(5).Hours().Tag("default").Do(func() {
		// 	fmt.Println(time.Now(), "scheduled")
		// })
	})

	return nil
}
