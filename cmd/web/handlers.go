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

func (app *application) certificates(w http.ResponseWriter, r *http.Request) {
	data := templateData{
		Certificates: content.Certificates,
	}
	app.render(w, r, http.StatusOK, "certificates.html", data)
}
