package httpio

import (
	"encoding/json"
	"net/http"
)

// RespondJSON ...
func RespondJSON(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(v)
}
