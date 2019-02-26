package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/mohuk/genie/manager"
)

// GetColumns ...
func GetColumns(manager manager.GenieManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		dbname := mux.Vars(r)["dbname"]
		tableName := mux.Vars(r)["tableId"]
		tf, err := manager.GetColumns(dbname, tableName)
		if err != nil {
			boom.Internal(w, err)
			return
		}
		err = json.NewEncoder(w).Encode(tf)
		if err != nil {
			boom.Internal(w, err)
			return
		}
	}
}
