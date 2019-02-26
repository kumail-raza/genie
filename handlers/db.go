package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/mohuk/genie/manager"
)

// GetDatabases ...
func GetDatabases(manager manager.GenieManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		dbs, err := manager.GetDatabases()
		if err != nil {
			boom.Internal(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(dbs)
		if err != nil {
			boom.Internal(w, err)
			return
		}
	}
}
