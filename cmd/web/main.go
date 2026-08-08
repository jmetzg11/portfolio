package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"portfolio.jmetzg11/internal/stocks"
)

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
	stocks        *stocks.Client
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	stocksURL := os.Getenv("STOCKS_API_URL")
	if stocksURL == "" {
		logger.Error("STOCKS_API_URL is not set")
		os.Exit(1)
	}

	app := &application{
		logger:        logger,
		templateCache: templateCache,
		stocks:        stocks.New(stocksURL),
	}

	srv := &http.Server{
		Addr:         ":4000",
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("starting server")

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
