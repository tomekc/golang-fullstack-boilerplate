package server

import (
	"fmt"
	"log"
	"net/http"

	"boilerplate/server/middleware"
	"boilerplate/server/views"
)

type Server struct {
	app              *Application
	globalMiddleware []middleware.Middleware
}

func NewServer(app *Application) *Server {
	return &Server{
		app: app,
	}
}

func (s *Server) Use(m middleware.Middleware) {
	s.globalMiddleware = append(s.globalMiddleware, m)
}

func (s *Server) Start() {
	rootMux := http.NewServeMux()

	s.Use(middleware.Logging)
	s.Use(middleware.CORS)

	s.app.RegisterRoutes(rootMux)
	views.RegisterStaticRoutes(rootMux)

	addr := fmt.Sprintf(":%d", s.app.Config.Server.Port)
	log.Printf("Server listening on %s\n", addr)

	if err := http.ListenAndServe(addr, middleware.Chain(rootMux, s.globalMiddleware...)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
