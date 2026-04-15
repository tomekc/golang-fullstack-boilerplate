package services

import (
	"io"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

type Client struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Company  string `db:"company"`
	City     string `db:"city"`
	Progress int    `db:"progress"`
	Created  string `db:"created"`
}

type ExampleService struct {
	db *sqlx.DB
}

func NewExampleService(db *sqlx.DB) *ExampleService {
	return &ExampleService{db: db}
}

func (s *ExampleService) GetClients() ([]Client, error) {
	var clients []Client
	err := s.db.Select(&clients, `SELECT id, name, company, city, progress, created FROM clients ORDER BY id`)
	return clients, err
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
