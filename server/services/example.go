package services

import (
	"fmt"
	"io"
	"net/http"
	"strings"
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

var allowedSortFields = map[string]bool{
	"name":     true,
	"progress": true,
}

type SortOrder struct {
	Field string
	Dir   string
}

func ParseSortOrder(s string) SortOrder {
	field, dir := "name", "asc"
	parts := strings.SplitN(s, ",", 2)
	if len(parts) >= 1 && allowedSortFields[parts[0]] {
		field = parts[0]
	}
	if len(parts) == 2 && parts[1] == "desc" {
		dir = "desc"
	}
	return SortOrder{Field: field, Dir: dir}
}

type ExampleService struct {
	db *sqlx.DB
}

func NewExampleService(db *sqlx.DB) *ExampleService {
	return &ExampleService{db: db}
}

func (s *ExampleService) GetClients(sort SortOrder) ([]Client, error) {
	var clients []Client
	dir := "ASC"
	if sort.Dir == "desc" {
		dir = "DESC"
	}
	query := fmt.Sprintf(`SELECT id, name, company, city, progress, created FROM clients ORDER BY %s %s`, sort.Field, dir)
	err := s.db.Select(&clients, query)
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
