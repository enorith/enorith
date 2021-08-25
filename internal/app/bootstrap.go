package app

import (
	"github.com/enorith/enorith/locales"
	"github.com/enorith/enorith/resources"
	"github.com/enorith/framework"
	"github.com/enorith/framework/cache"
	"github.com/enorith/framework/database"
	"github.com/enorith/framework/language"
	"github.com/enorith/http/view"
)

func BootstrapApp(app *framework.App) {
	app.Register(&database.Service{})
	app.Register(&cache.Service{})
	app.Register(language.NewService(locales.FS))
	view.WithDefault(resources.FS, "html", "views")
}
