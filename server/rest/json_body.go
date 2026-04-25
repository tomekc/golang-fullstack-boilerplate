package rest

import (
	"encoding/json"
	"net/http"
)

// Naive JSON body decoder. On production you may want to:
// check content type and handle errors, limit body size, etc.
func DecodeJSONBody(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
