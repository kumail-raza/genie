package handlers

import (
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/mohuk/genie/errors"
	"github.com/mohuk/genie/httpio"
	"github.com/mohuk/genie/manager"
)

// GetDatabases ...
func GetDatabases(manager manager.GenieManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbs, err := manager.GetDatabases()
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
		err = httpio.RespondJSON(w, dbs)
		if err != nil {
			boom.ExpectationFailed(w, err.Error())
			return
		}
	}
}
