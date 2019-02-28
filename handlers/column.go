package handlers

import (
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/mohuk/genie/errors"
	"github.com/mohuk/genie/httpio"
	"github.com/mohuk/genie/manager"
)

// GetColumns ...
func GetColumns(manager manager.GenieManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbname := mux.Vars(r)["dbname"]
		tableName := mux.Vars(r)["tableId"]
		tf, err := manager.GetColumns(dbname, tableName)
		if err != nil {
			switch err.(type) {
			case *errors.ErrDbConn:
				boom.ExpectationFailed(w, err.Error())
				return
			case *errors.ErrNoRows:
				boom.NotFound(w, err.Error())
				return
			default:
				boom.BadData(w, err.Error())
				return
			}
		}
		err = httpio.RespondJSON(w, tf)
		if err != nil {
			boom.ExpectationFailed(w, err.Error())
			return
		}
	}
}
