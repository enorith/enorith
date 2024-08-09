package app

import (
	"log"

	"github.com/enorith/enorith/internal/app/models"
	"github.com/enorith/enorith/internal/app/services"
	"github.com/enorith/enorith/internal/pkg/auth"
	"github.com/enorith/enorith/internal/pkg/path"
	"github.com/enorith/enorith/resources"
	"github.com/enorith/framework"
	"github.com/enorith/framework/authentication"
	"github.com/enorith/framework/cache"
	"github.com/enorith/framework/crond"
	"github.com/enorith/framework/database"
	"github.com/enorith/framework/http"
	"github.com/enorith/framework/http/session"
	"github.com/enorith/framework/language"
	"github.com/enorith/framework/queue"
	"github.com/enorith/framework/redis"
	"github.com/enorith/http/view"
	"gorm.io/gorm"
)

// BootstrapApp bootstrap application, register services
func BootstrapApp(app *framework.App, logDir string) {
	database.Migrator = Migration

	app.Register(database.NewService())
	app.Register(redis.Service{})

	app.Register(cache.Service{})
	language.Dir = "locales"
	app.Register(language.NewService(resources.FS, app.GetConfig().Locale))
	app.Register(authentication.NewAuthService())
	app.Register(auth.Service{})

	WithQueue(app)
	WithSchedule(app)
	WithHttp(app)
}

func WithQueue(app *framework.App) {
	app.Register(queue.NewService())
	app.Register(services.QueueService{})
}

func WithSchedule(app *framework.App) {
	app.Register(crond.Service{})
	// cron tasks
	app.Register(services.ScheduleService{})
}

func WithHttp(app *framework.App) {
	service := http.NewService()
	app.Register(services.HttpService{})
	app.Register(service)
	app.Register(session.NewService(path.BasePath("storage")))
	view.WithDefault(resources.FS, "html", "views")
}

func Migration(tx *gorm.DB) {
	log.Print("migration.....")
	tx.Migrator().AutoMigrate(&models.User{})
}
