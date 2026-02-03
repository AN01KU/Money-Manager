package service

import (
	"errors"
)

var UnauthorizedError = errors.New("Invalid token")

// func Authorization(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var authToken = r.Header.Get("Authorization")
// 		var token = strings.TrimPrefix(authToken, "Bearer ")
// 		var err error

// 		if token == "" {
// 			api.RequestErrorHandler(w, UnauthorizedError)
// 			return
// 		}

// 		// var database *tools.DatabaseInterface
// 		// database, err = tools.NewDatabase()
// 		// if err != nil {
// 		// 	api.InternalErrorHandler(w)
// 		// 	return
// 		// }

// 		// var userDetails *tools.UserDetails
// 		// userDetails = database.

// 		next.ServeHTTP(w, r)
// 	})
// }
