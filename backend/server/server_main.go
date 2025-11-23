package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	ListenPort int
	apiMux     *http.ServeMux
}

func New(port int) *Server {
	return &Server{
		ListenPort: port,
		apiMux:     http.NewServeMux(),
	}
}

func (s *Server) Start() {
	// Create root router
	rootMux := http.NewServeMux()
	rootMux.Handle("/api/", http.StripPrefix("/api", s.apiMux))

	addr := fmt.Sprintf(":%d", s.ListenPort)
	fmt.Printf("Server listening on %s\n", addr)

	if err := http.ListenAndServe(addr, rootMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
