package services

import (
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/enorith/framework"
	"github.com/enorith/framework/crond"
	"github.com/enorith/supports/carbon"
	"github.com/go-co-op/gocron"
)

type ScheduleService struct {
}

// Register service when app starting, before http server start
// you can configure service, prepare global vars etc.
// running at main goroutine
func (s ScheduleService) Register(app *framework.App) error {

	app.Resolving(func(s *gocron.Scheduler /*, tx *gorm.DB // inject your dependencies*/) {
		// run tasks
		// s.Every(5).Hours().Tag("default").Do(func() {
		// 	fmt.Println(time.Now(), "scheduled")
		// })
	})
	app.Daemon(func(exit chan struct{}) {
		time.Sleep(time.Second)
		printSchedules(crond.Scheduler)
		<-exit
	})

	return nil
}

func printSchedules(s *gocron.Scheduler) {

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Tags"},
			{Align: simpletable.AlignCenter, Text: "Next"},
			{Align: simpletable.AlignCenter, Text: "Last"},
		},
	}
	for _, j := range s.Jobs() {
		table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: strings.Join(j.Tags(), ",")},
			{Align: simpletable.AlignLeft, Text: j.NextRun().Format(carbon.DefaultDateTimeFormat)},
			{Align: simpletable.AlignLeft, Text: j.LastRun().Format(carbon.DefaultDateTimeFormat)},
		})
	}
	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Cron schedules", Span: 3},
		},
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
