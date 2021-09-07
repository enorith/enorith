package app

import (
	"log"

	"github.com/enorith/enorith/internal/app/models"
	"github.com/enorith/enorith/internal/app/routes"
	"github.com/enorith/enorith/internal/pkg/auth"
	"github.com/enorith/enorith/internal/pkg/env"
	"github.com/enorith/enorith/locales"
	"github.com/enorith/enorith/resources"
	"github.com/enorith/framework"
	"github.com/enorith/framework/authentication"
	"github.com/enorith/framework/cache"
	"github.com/enorith/framework/database"
	"github.com/enorith/framework/http"
	"github.com/enorith/framework/language"
	"github.com/enorith/framework/queue"
	"github.com/enorith/framework/redis"
	h "github.com/enorith/http"
	"github.com/enorith/http/router"
	"github.com/enorith/http/view"
	"gorm.io/gorm"
)

func BootstrapApp(app *framework.App) {
	database.Migrator = Migration

	app.Register(env.Service{})
	app.Register(database.NewService())
	app.Register(cache.Service{})
	app.Register(redis.Service{})
	app.Register(language.NewService(locales.FS, app.GetConfig().Locale))
	app.Register(authentication.NewAuthService())
	app.Register(auth.Service{})
	app.Register(queue.NewService())

	WithHttp(app)
	view.WithDefault(resources.FS, "html", "views")
}

func WithHttp(app *framework.App) {
	service := http.NewService(func(rw *router.Wrapper, k *h.Kernel) {
		k.KeepAlive()
		k.OutputLog = true
		RegisterRoutes(rw)
	})

	app.Register(service)
}

func RegisterRoutes(rw *router.Wrapper) {
	routes.WebRoutes(rw)

	rw.Group(func(r *router.Wrapper) {
		routes.ApiRoutes(r)
	}, "api")

}

func Migration(tx *gorm.DB) {
	log.Print("migration.....")
	tx.Migrator().AutoMigrate(&models.User{})
}
