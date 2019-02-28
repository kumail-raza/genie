package handlers

import (
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/mohuk/genie/errors"
	"github.com/mohuk/genie/httpio"
	"github.com/mohuk/genie/manager"
)

// GetTables ...
func GetTables(manager manager.GenieManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbname := mux.Vars(r)["dbname"]
		tables, err := manager.GetTables(dbname)
		if err != nil {
			switch err.(type) {
			case *errors.ErrDbConn:
				boom.ExpectationFailed(w, err.Error())
				return
			default:
				boom.BadData(w, err.Error())
				return
			}
		}
		err = httpio.RespondJSON(w, tables)
		if err != nil {
			boom.ExpectationFailed(w, err.Error())
			return
		}
	}
}
