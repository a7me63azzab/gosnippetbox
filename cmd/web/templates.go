package main

import (
	"html/template"
	"path/filepath"
	"time"

	"snippetbox.azzab.net/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet
	Form        any
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.go.tmpl")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		fileName := filepath.Base(page)

		ts, err := template.New(fileName).Funcs(functions).ParseFiles("./ui/html/base.go.tmpl")

		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.go.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)

		if err != nil {
			return nil, err
		}

		cache[fileName] = ts
	}

	return cache, nil
}
