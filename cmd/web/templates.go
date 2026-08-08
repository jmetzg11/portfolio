package main

import (
	"html/template"
	"io/fs"
	"path/filepath"

	"portfolio.jmetzg11/internal/content"
	"portfolio.jmetzg11/internal/stocks"
	"portfolio.jmetzg11/ui"
)

type templateData struct {
	IsHome        bool
	BackURL       string
	BackLabel     string
	Projects      []content.Project
	GithubIcon    string
	BJJTechniques []content.BJJTechnique
	Certificates  []content.Certificate
	Years         []content.YearOption
	MapData       template.JS
	Stocks        []stocks.DisplayStock
	Crypto        []stocks.DisplayCoin
	CDMXSections  []content.CDMXSection
	CDMXToBook    []content.CDMXSection
	CDMXDetail    content.CDMXDetail
	CDMXPlace     content.CDMXPlace
	CDMXText      content.CDMXStrings
	Lang          string
}

func (d templateData) IsRU() bool {
	return d.Lang == content.LangRU
}

// LangQuery rides along on every internal link, so the language survives a click
// into a detail page and the click back out.
func (d templateData) LangQuery() string {
	if d.Lang == content.LangRU {
		return "?lang=" + content.LangRU
	}
	return ""
}

func (d templateData) BackURLOrDefault() string {
	if d.BackURL == "" {
		return "/"
	}
	return d.BackURL
}

func (d templateData) BackLabelOrDefault() string {
	if d.BackLabel == "" {
		return "Home"
	}
	return d.BackLabel
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
