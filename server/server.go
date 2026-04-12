package server

import (
	"fmt"
	"log"
	"net/http"

	"boilerplate/server/config"
	"boilerplate/server/middleware"
	"boilerplate/server/views"
)

type Server struct {
	ListenPort       int
	apiMux           *http.ServeMux
	globalMiddleware []middleware.Middleware
}

func New(cfg config.Server) *Server {
	return &Server{
		ListenPort: cfg.Port,
		apiMux:     http.NewServeMux(),
	}
}

func (s *Server) Use(m middleware.Middleware) {
	s.globalMiddleware = append(s.globalMiddleware, m)
}

func (s *Server) Start() {
	// Create root router
	rootMux := http.NewServeMux()
	decoratedMux := middleware.Chain(s.apiMux, s.globalMiddleware...)

	// API handlers
	rootMux.Handle("/api/", http.StripPrefix("/api", decoratedMux))

	// Serve frontend
	log.Println("Frontend: Templ + HTMX (server-side rendering)")
	views.RegisterPageRoutes(rootMux)
	views.RegisterStaticRoutes(rootMux)

	addr := fmt.Sprintf(":%d", s.ListenPort)
	log.Printf("Server listening on %s\n", addr)

	if err := http.ListenAndServe(addr, rootMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
