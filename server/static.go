package server

import (
	"io/fs"
	"log"
	"net/http"
)

func (s *Server) serveStatic(mux *http.ServeMux) {
	staticFS, err := fs.Sub(s.staticFS, "build-frontend")
	if err != nil {
		log.Fatalf("Failed to create sub filesystem: %v", err)
	}

	fileServer := http.FileServer(http.FS(staticFS))
	mux.Handle("/", fileServer)
}
