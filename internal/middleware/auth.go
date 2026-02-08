package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/AN01KU/money-manager/internal/api"
	"github.com/AN01KU/money-manager/internal/tools"
)

type contextKey string

const UserContextKey contextKey = "user"

func Auth(db tools.DatabaseInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				api.UnauthorizedErrorHandler(w)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				api.UnauthorizedErrorHandler(w)
				return
			}

			userID, err := tools.ValidateJWTToken(tokenString)
			if err != nil {
				api.RequestErrorHandler(w, err)
				return
			}

			user := db.GetUserByID(userID)
			if user == nil {
				api.RequestErrorHandler(w, errors.New("user not found"))
				return
			}
			ctx := context.WithValue(r.Context(), UserContextKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUser(ctx context.Context) (*tools.User, bool) {
	user, ok := ctx.Value(UserContextKey).(*tools.User)
	return user, ok
}
