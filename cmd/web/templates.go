package main

import (
	"html/template"
	"io/fs"
	"path/filepath"

	"portfolio.jmetzg11/internal/content"
	"portfolio.jmetzg11/ui"
)

type templateData struct {
	Projects      []content.Project
	GithubIcon    string
	BJJTechniques []content.BJJTechnique
	Certificates  []content.Certificate
	Years         []content.YearOption
	MapData       template.JS
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		patterns := []string{
			"html/base.html",
			page,
		}

		ts, err := template.New(name).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
