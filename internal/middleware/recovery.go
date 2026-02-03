package middleware

import (
	"net/http"

	"github.com/AN01KU/money-manager/internal/api"
	log "github.com/sirupsen/logrus"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.WithField("panic", err).Error("recovered from panic")
				api.WriteJSON(w, 500, api.Error{
					Code:    500,
					Message: "Internal server error",
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}
