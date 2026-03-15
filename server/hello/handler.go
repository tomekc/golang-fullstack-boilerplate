package hello

import (
	"encoding/json"
	"net/http"
	"time"

	"boilerplate/server/rest"
)

type HelloResponse struct {
	Time string `json:"time"`
}

func GetTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := HelloResponse{
		Time: time.Now().Format(time.RFC3339),
	}

	rest.Ok(w, response)
}

func Echo(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	rest.Ok(w, map[string]json.RawMessage{"body": body})
}
