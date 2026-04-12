package middleware

import "net/http"

type flightRecorder struct {
	http.ResponseWriter
	Status       int
	BytesWritten int
}

func newFlightRecorder(w http.ResponseWriter) *flightRecorder {
	return &flightRecorder{ResponseWriter: w, Status: http.StatusOK}
}

func (r *flightRecorder) WriteHeader(code int) {
	r.Status = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *flightRecorder) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	r.BytesWritten += n
	return n, err
}
