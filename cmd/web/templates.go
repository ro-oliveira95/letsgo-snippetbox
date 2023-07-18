package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/ro-oliveira95/letsgo-snippetbox/internal/models"
	"github.com/ro-oliveira95/letsgo-snippetbox/ui"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

// Create a humanDate function which returns a nicely formatted string
// representation of a time.Time object.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// String-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var tFuncs = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}

		ts, err := template.New(name).Funcs(tFuncs).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
