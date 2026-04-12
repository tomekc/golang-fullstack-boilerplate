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
	globalMiddleware []middleware.Middleware
}

func New(cfg config.Server) *Server {
	return &Server{
		ListenPort: cfg.Port,
	}
}

func (s *Server) Use(m middleware.Middleware) {
	s.globalMiddleware = append(s.globalMiddleware, m)
}

func (s *Server) Start() {
	rootMux := http.NewServeMux()

	s.Use(middleware.Logging)
	s.Use(middleware.CORS)

	views.RegisterPageRoutes(rootMux)
	views.RegisterStaticRoutes(rootMux)

	addr := fmt.Sprintf(":%d", s.ListenPort)
	log.Printf("Server listening on %s\n", addr)

	if err := http.ListenAndServe(addr, middleware.Chain(rootMux, s.globalMiddleware...)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
