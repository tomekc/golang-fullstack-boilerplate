package views

import (
	"net/http"

	"boilerplate/server/services"
	"boilerplate/server/views/fragments"
	"boilerplate/server/views/pages"
)

type Handlers struct {
	ExampleService *services.ExampleService
}

func NewHandlers(exampleService *services.ExampleService) *Handlers {
	return &Handlers{
		ExampleService: exampleService,
	}
}

func (h *Handlers) DashboardPage(w http.ResponseWriter, r *http.Request) {
	pages.Dashboard().Render(r.Context(), w)
}

func (h *Handlers) HelloPage(w http.ResponseWriter, r *http.Request) {
	pages.Hello().Render(r.Context(), w)
}

func (h *Handlers) AboutPage(w http.ResponseWriter, r *http.Request) {
	pages.About().Render(r.Context(), w)
}

func (h *Handlers) HtmxGetTime(w http.ResponseWriter, r *http.Request) {
	t := h.ExampleService.GetTime()
	fragments.TimeResult(t).Render(r.Context(), w)
}

func (h *Handlers) HtmxEcho(w http.ResponseWriter, r *http.Request) {
	info := h.ExampleService.BuildEchoResponse(r)
	fragments.EchoResult(info).Render(r.Context(), w)
}
