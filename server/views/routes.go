package views

import "net/http"

func RegisterPageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", DashboardPage)
	mux.HandleFunc("GET /hello", HelloPage)
	mux.HandleFunc("GET /about", AboutPage)
	mux.HandleFunc("GET /htmx/time", HtmxGetTime)
	mux.HandleFunc("POST /htmx/echo", HtmxEcho)
}
