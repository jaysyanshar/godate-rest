package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jaysyanshar/godate-rest/config"
)

func TestJWTMiddleware(t *testing.T) {
	cfg := &config.Config{
		JwtSecret: "godate-auth-secret",
	}
	middleware := NewMiddleware(cfg)

	tests := []struct {
		name           string
		token          string
		expectedStatus int
	}{
		{
			name:           "valid token",
			token:          "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjM1NTkxMzUsImlzcyI6IkdvRGF0ZSIsInN1YiI6IjEifQ.yvPhg8skGWTFA6rq7NClHbNOE2uFmNXQZkfbtY0R0DA",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "missing token",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "invalid token",
			token:          "invalid_token_here",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "invalid token with Bearer prefix",
			token:          "Bearer invalid_token_here",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP request
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			// Add the token to the request header if it's provided
			if tt.token != "" {
				req.Header.Set("X-Authorization", tt.token)
			}

			// Create a new HTTP response recorder
			res := httptest.NewRecorder()

			// Call the JWTMiddleware method
			middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})).ServeHTTP(res, req)

			// Check the response status code
			if res.Code != tt.expectedStatus {
				t.Errorf("expected status code %d, got %d", tt.expectedStatus, res.Code)
			}
		})
	}
}
