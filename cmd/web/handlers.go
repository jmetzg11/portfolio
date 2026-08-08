package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"portfolio.jmetzg11/internal/content"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := templateData{
		IsHome:     true,
		Projects:   content.Projects,
		GithubIcon: content.GithubIconPath,
	}
	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) investing(w http.ResponseWriter, r *http.Request) {
	summary, err := app.stocks.Summary()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	display := summary.Display()

	data := templateData{
		Stocks: display.Stocks,
		Crypto: display.Crypto,
	}
	app.render(w, r, http.StatusOK, "investing.html", data)
}

func (app *application) bjj(w http.ResponseWriter, r *http.Request) {
	data := templateData{
		BJJTechniques: content.BJJTechniques,
	}
	app.render(w, r, http.StatusOK, "bjj.html", data)
}

func (app *application) geography(w http.ResponseWriter, r *http.Request) {
	encoded, err := json.Marshal(content.MarkersByYear)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Years:   content.Years(),
		MapData: template.JS(encoded),
	}
	app.render(w, r, http.StatusOK, "geography.html", data)
}

func (app *application) cdmx(w http.ResponseWriter, r *http.Request) {
	lang := content.CDMXLang(r.URL.Query().Get("lang"))

	encoded, err := json.Marshal(content.CDMXMarkers(lang))
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		CDMXSections: content.CDMXSections(lang),
		CDMXToBook:   content.CDMXToBook(lang),
		CDMXText:     content.CDMXText(lang),
		Lang:         lang,
		MapData:      template.JS(encoded),
	}
	app.render(w, r, http.StatusOK, "cdmx.html", data)
}

func (app *application) cdmxDetail(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	lang := content.CDMXLang(r.URL.Query().Get("lang"))

	detail, ok := content.CDMXDetailBySlug(slug, lang)
	if !ok {
		http.NotFound(w, r)
		return
	}

	place, ok := content.CDMXPlaceBySlug(slug, lang)
	if !ok {
		http.NotFound(w, r)
		return
	}

	data := templateData{
		CDMXDetail: detail,
		CDMXPlace:  place,
		CDMXText:   content.CDMXText(lang),
		Lang:       lang,
		BackLabel:  "CDMX",
	}
	data.BackURL = "/cdmx" + data.LangQuery()

	app.render(w, r, http.StatusOK, "cdmx_detail.html", data)
}

func (app *application) farmsville(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "farmsville.html", templateData{})
}

func (app *application) certificates(w http.ResponseWriter, r *http.Request) {
	data := templateData{
		Certificates: content.Certificates,
	}
	app.render(w, r, http.StatusOK, "certificates.html", data)
}
