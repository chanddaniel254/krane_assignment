package middleware

import (
	"context"
	"net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check headers or perform authentication logic here
		userHeader := r.Header.Get("Token")

		// Add more authentication and authorization checks as needed
		ctx := context.WithValue(r.Context(), "Token", userHeader)

		// If authentication and authorization checks pass, continue to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
