package server

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"boilerplate/server/config"
	"boilerplate/server/middleware"
	"boilerplate/server/views"
)

type Server struct {
	ListenPort       int
	frontend         config.FrontendStack
	apiMux           *http.ServeMux
	staticFS         embed.FS
	globalMiddleware []middleware.Middleware
}

func New(cfg config.Server, staticFS embed.FS) *Server {
	return &Server{
		ListenPort: cfg.Port,
		frontend:   cfg.Frontend,
		apiMux:     http.NewServeMux(),
		staticFS:   staticFS,
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
	if s.frontend == config.FrontendTempl {
		log.Println("Frontend: Templ + HTMX (server-side rendering)")
		views.RegisterPageRoutes(rootMux)
		views.RegisterStaticRoutes(rootMux)
	} else {
		log.Println("Frontend: Svelte SPA (client-side rendering)")
		s.serveStatic(rootMux)
	}

	addr := fmt.Sprintf(":%d", s.ListenPort)
	log.Printf("Server listening on %s\n", addr)

	if err := http.ListenAndServe(addr, rootMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
