package views

import "net/http"

func RegisterPageRoutes(mux *http.ServeMux, h *Handlers) {
	mux.HandleFunc("GET /{$}", h.DashboardPage)
	mux.HandleFunc("GET /hello", h.HelloPage)
	mux.HandleFunc("GET /about", h.AboutPage)
	mux.HandleFunc("GET /htmx/time", h.HtmxGetTime)
	mux.HandleFunc("POST /htmx/echo", h.HtmxEcho)
}
