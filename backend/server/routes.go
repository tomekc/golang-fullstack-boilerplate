package server

import (
	"net/http"
)

func (s *Server) ApiEndpoint(pattern string, handler http.HandlerFunc) {
	s.apiMux.HandleFunc(pattern, handler)
}
