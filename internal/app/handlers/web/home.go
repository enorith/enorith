package web

import (
	"github.com/enorith/http/content"
	"github.com/enorith/http/view"
)

func Index() (*content.TemplateResponse, error) {
	return view.View("welcome", 200, map[string]string{
		"Title": "Enorith",
		"Desc":  "A framework for golang web artisan",
	})
}
