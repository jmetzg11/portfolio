package main

import (
	"net/http"

	"portfolio.jmetzg11/internal/content"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := templateData{
		Projects:   content.Projects,
		GithubIcon: content.GithubIconPath,
	}
	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) bjj(w http.ResponseWriter, r *http.Request) {
	data := templateData{
		BJJTechniques: content.BJJTechniques,
	}
	app.render(w, r, http.StatusOK, "bjj.html", data)
}

func (app *application) certificates(w http.ResponseWriter, r *http.Request) {
	data := templateData{
		Certificates: content.Certificates,
	}
	app.render(w, r, http.StatusOK, "certificates.html", data)
}
