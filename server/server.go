package server

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"boilerplate/server/config"
	"boilerplate/server/middleware"
)

type Server struct {
	ListenPort       int
	apiMux           *http.ServeMux
	staticFS         embed.FS
	globalMiddleware []middleware.Middleware
}

func New(cfg config.Server, staticFS embed.FS, middlewares ...middleware.Middleware) *Server {
	return &Server{
		ListenPort:       cfg.Port,
		apiMux:           http.NewServeMux(),
		staticFS:         staticFS,
		globalMiddleware: middlewares,
	}
}

func (s *Server) Start() {
	// Create root router
	rootMux := http.NewServeMux()
	decoratedMux := middleware.Chain(s.apiMux, s.globalMiddleware...)

	// API handlers
	rootMux.Handle("/api/", http.StripPrefix("/api", decoratedMux))

	// Serve static files
	s.serveStatic(rootMux)

	addr := fmt.Sprintf(":%d", s.ListenPort)
	log.Printf("Server listening on %s\n", addr)

	if err := http.ListenAndServe(addr, rootMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
