package server

import (
	"net/http"

	"boilerplate/server/services"
	"boilerplate/server/views/fragments"
	"boilerplate/server/views/pages"
)

func (app *Application) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
	})
	mux.HandleFunc("GET /dashboard", app.DashboardPage)
	mux.HandleFunc("GET /hello", app.HelloPage)
	mux.HandleFunc("GET /about", app.AboutPage)
	mux.HandleFunc("GET /htmx/clients", app.HtmxClients)
	mux.HandleFunc("GET /htmx/time", app.HtmxGetTime)
	mux.HandleFunc("POST /htmx/echo", app.HtmxEcho)
}

func (app *Application) DashboardPage(w http.ResponseWriter, r *http.Request) {
	sort := services.ParseSortOrder(r.URL.Query().Get("sort"))
	clients, err := app.ExampleService.GetClients(sort)
	if err != nil {
		http.Error(w, "failed to load clients", http.StatusInternalServerError)
		return
	}
	pages.Dashboard(clients, sort).Render(r.Context(), w)
}

func (app *Application) HtmxClients(w http.ResponseWriter, r *http.Request) {
	sort := services.ParseSortOrder(r.URL.Query().Get("sort"))
	clients, err := app.ExampleService.GetClients(sort)
	if err != nil {
		http.Error(w, "failed to load clients", http.StatusInternalServerError)
		return
	}
	fragments.ClientsTable(clients, sort).Render(r.Context(), w)
}

func (app *Application) HelloPage(w http.ResponseWriter, r *http.Request) {
	pages.Hello().Render(r.Context(), w)
}

func (app *Application) AboutPage(w http.ResponseWriter, r *http.Request) {
	pages.About().Render(r.Context(), w)
}

func (app *Application) HtmxGetTime(w http.ResponseWriter, r *http.Request) {
	t := app.ExampleService.GetTime()
	fragments.TimeResult(t).Render(r.Context(), w)
}

func (app *Application) HtmxEcho(w http.ResponseWriter, r *http.Request) {
	info := app.ExampleService.BuildEchoResponse(r)
	fragments.EchoResult(info).Render(r.Context(), w)
}
