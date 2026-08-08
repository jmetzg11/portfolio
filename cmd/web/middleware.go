package main

import (
	"fmt"
	"net/http"
)

// csp whitelists the only external origins the site uses: unpkg for Leaflet,
// OpenStreetMap and Carto for map tiles, and YouTube for the BJJ thumbnails.
// style-src allows inline because the geography year buttons carry a
// style attribute; injected CSS is a far weaker attack than injected script,
// which stays strict.
const csp = "default-src 'self'; " +
	"script-src 'self' https://unpkg.com; " +
	"style-src 'self' 'unsafe-inline' https://unpkg.com; " +
	"img-src 'self' https://unpkg.com https://*.tile.openstreetmap.org https://*.basemaps.cartocdn.com https://img.youtube.com; " +
	"frame-ancestors 'none'; " +
	"base-uri 'self'"

func commonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", csp)
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			pv := recover()
			if pv != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, r, fmt.Errorf("%v", pv))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
