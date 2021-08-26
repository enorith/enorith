package app

import (
	"log"

	"github.com/enorith/enorith/internal/app/models"
	"github.com/enorith/enorith/locales"
	"github.com/enorith/enorith/resources"
	"github.com/enorith/framework"
	"github.com/enorith/framework/cache"
	"github.com/enorith/framework/database"
	"github.com/enorith/framework/language"
	"github.com/enorith/framework/redis"
	"github.com/enorith/http/view"
	"gorm.io/gorm"
)

func BootstrapApp(app *framework.App) {
	database.Migrator = Migration

	app.Register(&database.Service{})
	app.Register(&cache.Service{})
	app.Register(&redis.Service{})
	app.Register(language.NewService(locales.FS, app.GetConfig().Locale))
	view.WithDefault(resources.FS, "html", "views")
}

func Migration(tx *gorm.DB) {
	log.Print("migration.....")
	tx.Migrator().AutoMigrate(&models.User{})
}
