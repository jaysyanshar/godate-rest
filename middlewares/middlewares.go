package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/jaysyanshar/godate-rest/config"
)

var cfg = config.Get()

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the JWT token from the request header
		tokenString := r.Header.Get("Authorization")

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Provide the secret key or public key to verify the token
			return []byte(cfg.JwtSecret), nil
		})

		// Handle token parsing errors
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		// Check if the token is valid
		if token.Valid {
			// Token is valid, call the next handler
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
	})
}
