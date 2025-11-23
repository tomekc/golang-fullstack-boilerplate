package hello

import (
	"net/http"
	"time"

	"boilerplate/backend/rest"
)

type HelloResponse struct {
	Time string `json:"time"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := HelloResponse{
		Time: time.Now().Format(time.RFC3339),
	}

	rest.Ok(w, response)
}
