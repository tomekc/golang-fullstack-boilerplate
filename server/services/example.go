package services

import (
	"io"
	"net/http"
	"time"
)

type ExampleService struct{}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}

func (s *ExampleService) GetTime() string {
	return time.Now().Format(time.RFC3339)
}

func (s *ExampleService) BuildEchoResponse(r *http.Request) map[string]string {
	body, _ := io.ReadAll(r.Body)
	return map[string]string{
		"Method":      r.Method,
		"Remote Addr": r.RemoteAddr,
		"User-Agent":  r.UserAgent(),
		"Host":        r.Host,
		"HX-Request":  r.Header.Get("HX-Request"),
		"HX-Target":   r.Header.Get("HX-Target"),
		"Timestamp":   time.Now().Format(time.RFC3339),
		"Body":        string(body),
	}
}
