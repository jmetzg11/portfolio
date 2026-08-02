package main

import (
	"net/http"

	"portfolio.jmetzg11/ui"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("GET /ping", ping)

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /bjj", app.bjj)
	return app.recoverPanic(commonHeaders(mux))
}
