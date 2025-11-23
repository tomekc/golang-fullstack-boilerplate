package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	ListenPort       int
	apiMux           *http.ServeMux
	globalMiddleware []Middleware
}

func New(port int, middlewares ...Middleware) *Server {
	return &Server{
		ListenPort:       port,
		apiMux:           http.NewServeMux(),
		globalMiddleware: middlewares,
	}
}

func (s *Server) Start() {
	// Create root router
	rootMux := http.NewServeMux()
	decoratedMux := Chain(s.apiMux, s.globalMiddleware...)
	rootMux.Handle("/api/", http.StripPrefix("/api", decoratedMux))

	addr := fmt.Sprintf(":%d", s.ListenPort)
	fmt.Printf("Server listening on %s\n", addr)

	if err := http.ListenAndServe(addr, rootMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
