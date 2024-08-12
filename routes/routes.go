// routes/routes.go
package routes

import (
	"github.com/gorilla/mux"
	"github.com/jaysyanshar/godate-rest/controllers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HelloHandler).Methods("GET")
	return router
}
