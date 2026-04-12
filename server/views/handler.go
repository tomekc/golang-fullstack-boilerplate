package views

import (
	"io"
	"net/http"
	"time"

	"boilerplate/server/views/fragments"
	"boilerplate/server/views/pages"
)

func DashboardPage(w http.ResponseWriter, r *http.Request) {
	pages.Dashboard().Render(r.Context(), w)
}

func HelloPage(w http.ResponseWriter, r *http.Request) {
	pages.Hello().Render(r.Context(), w)
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	pages.About().Render(r.Context(), w)
}

func HtmxGetTime(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC3339)
	fragments.TimeResult(t).Render(r.Context(), w)
}

func HtmxEcho(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	info := map[string]string{
		"Method":      r.Method,
		"Remote Addr": r.RemoteAddr,
		"User-Agent":  r.UserAgent(),
		"Host":        r.Host,
		"HX-Request":  r.Header.Get("HX-Request"),
		"HX-Target":   r.Header.Get("HX-Target"),
		"Timestamp":   time.Now().Format(time.RFC3339),
		"Body":        string(body),
	}
	fragments.EchoResult(info).Render(r.Context(), w)
}
