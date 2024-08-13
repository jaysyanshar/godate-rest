// routes/routes.go
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaysyanshar/godate-rest/controllers/auth"
	"github.com/jaysyanshar/godate-rest/controllers/dashboard"
	"github.com/jaysyanshar/godate-rest/middlewares"
)

type router struct {
	Mid       middlewares.Middleware
	Auth      auth.AuthController
	Dashboard dashboard.DashboardController
}

func SetupRouter(mid middlewares.Middleware, auth auth.AuthController, dashboard dashboard.DashboardController) *mux.Router {
	r := router{
		Mid:       mid,
		Auth:      auth,
		Dashboard: dashboard,
	}
	router := mux.NewRouter()

	// dashboard
	router.HandleFunc("/", r.withMiddleware(r.Dashboard.HelloHandler, r.Mid.JWTMiddleware)).Methods("GET")

	// auth
	router.HandleFunc("/signup", r.Auth.SignUpHandler).Methods("POST")
	router.HandleFunc("/login", r.Auth.LoginHandler).Methods("POST")
	return router
}

// Helper function to wrap a handler with middleware
func (r router) withMiddleware(h http.HandlerFunc, middleware func(http.Handler) http.Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		middleware(http.HandlerFunc(h)).ServeHTTP(w, req)
	}
}
