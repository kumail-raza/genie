package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/mohuk/genie/manager"
)

// GetTables ...
func GetTables(manager manager.GenieManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		dbname := mux.Vars(r)["dbname"]
		tables, err := manager.GetTables(dbname)
		if err != nil {
			boom.Internal(w, err)
			return
		}
		err = json.NewEncoder(w).Encode(tables)
		if err != nil {
			boom.Internal(w, err)
			return
		}
	}
}
