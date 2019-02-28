package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func WithLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
