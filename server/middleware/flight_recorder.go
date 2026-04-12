package middleware

import "net/http"

// flightRecorder is a response wrapper that captures the HTTP status code.
// By default, http.ResponseWriter doesn't expose the status code after it's been written.
// flightRecorder embeds http.ResponseWriter and overrides WriteHeader to intercept
// and store the status code, enabling Logging to log it.
type flightRecorder struct {
	http.ResponseWriter
	Status int
}

// WriteHeader intercepts the status code, stores it in Status, and passes it through to the real writer.
func (r *flightRecorder) WriteHeader(code int) {
	r.Status = code
	r.ResponseWriter.WriteHeader(code)
}
