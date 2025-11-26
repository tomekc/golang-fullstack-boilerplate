package server

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"boilerplate/backend/config"
)

type Server struct {
	ListenPort       int
	apiMux           *http.ServeMux
	staticFS         embed.FS
	globalMiddleware []Middleware
}

func New(server config.Server, staticFS embed.FS, middlewares ...Middleware) *Server {
	return &Server{
		ListenPort:       server.Port,
		apiMux:           http.NewServeMux(),
		staticFS:         staticFS,
		globalMiddleware: middlewares,
	}
}

func (s *Server) Start() {
	// Create root router
	rootMux := http.NewServeMux()
	decoratedMux := Chain(s.apiMux, s.globalMiddleware...)
	rootMux.Handle("/api/", http.StripPrefix("/api", decoratedMux))

	// Serve static files
	s.serveStatic(rootMux)

	addr := fmt.Sprintf(":%d", s.ListenPort)
	log.Printf("Server listening on %s\n", addr)

	if err := http.ListenAndServe(addr, rootMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
