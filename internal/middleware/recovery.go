package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/AN01KU/money-manager/internal/api"
	log "github.com/sirupsen/logrus"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.WithField("panic", err).Error("recovered from panic")
				writeJSON(w, 500, api.Error{
					Code:    500,
					Message: "Internal server error",
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}
