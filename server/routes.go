package server

import (
	"context"
	"net/http"

	"boilerplate/server/services"
	"boilerplate/server/views/fragments"
	"boilerplate/server/views/pages"
	"go.jetify.com/sse"
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
	mux.HandleFunc("PUT /htmx/hello/remember", app.HtmxRemember)
	mux.HandleFunc("GET /example/sse/sales", app.ExampleSSESalesStream) // EXAMPLE: SSE — remove with example
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

func (app *Application) HtmxRemember(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	r.ParseForm()
	shouldRemember := r.PostForm.Get("checked")
	if shouldRemember == "true" {
		w.Write([]byte("<p>Remembered </p>"))
	} else {
		w.Write([]byte("<p>Not remembered </p>"))
	}
}

func (app *Application) ExampleSSESalesStream(w http.ResponseWriter, r *http.Request) {
	conn, err := sse.Upgrade(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	ch, unsubscribe := app.ExampleSSESales.Subscribe()
	defer unsubscribe()

	ctx := r.Context()
	if err := app.sendSalesEvent(conn, ctx, app.ExampleSSESales.Current()); err != nil {
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-ch:
			if !ok {
				return
			}
			if err := app.sendSalesEvent(conn, ctx, v); err != nil {
				return
			}
		}
	}
}

func (app *Application) sendSalesEvent(conn *sse.Conn, ctx context.Context, currentSales string) error {
	return conn.SendEvent(ctx, &sse.Event{
		Event: "sales",
		Data:  sse.Raw(currentSales),
	})
}
