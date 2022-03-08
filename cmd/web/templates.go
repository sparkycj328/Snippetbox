package main

import (
	"html/template"
	"path/filepath"

	"github.com/sparkycj328/Snippetbox/pkg/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

// newTemplateCache will create a slice of all filepaths associated with the page.tmpl string.
// It will then loop through the pages one by one and add the templates to the template set
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
