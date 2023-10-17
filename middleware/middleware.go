package middleware

import (
	"context"
	"net/http"

	"github.com/rs/cors"
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

func CorsMiddleware(next http.Handler) http.Handler {
	// Create a CORS middleware with your desired options
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Change this to allow specific origins
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	return corsHandler.Handler(next)
}
