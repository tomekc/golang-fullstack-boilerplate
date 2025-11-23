package server

import "net/http"

type flightRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *flightRecorder) WriteHeader(code int) {
	r.Status = code
	r.ResponseWriter.WriteHeader(code)
}
