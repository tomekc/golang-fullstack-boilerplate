package views

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static
var staticFiles embed.FS

func RegisterStaticRoutes(mux *http.ServeMux) {
	sub, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatalf("Failed to create static sub filesystem: %v", err)
	}
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(sub))))
}
