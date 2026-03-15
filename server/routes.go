package server

import (
	"net/http"

	"boilerplate/server/hello"
)

func (s *Server) ApiEndpoint(pattern string, handler http.HandlerFunc) {
	s.apiMux.HandleFunc(pattern, handler)
}

// Add our api endpoints here. Feel free to remove "/hello" if you don't need it.
func (s *Server) AddAPIRoutes() {
	s.ApiEndpoint("GET /hello", hello.GetTime)
	s.ApiEndpoint("POST /echo", hello.Echo)
}
