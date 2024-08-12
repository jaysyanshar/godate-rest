package main

import (
	"log"
	"net/http"

	"github.com/jaysyanshar/godate-rest/config"
	authCtrl "github.com/jaysyanshar/godate-rest/controllers/auth"
	dashboardCtrl "github.com/jaysyanshar/godate-rest/controllers/dashboard"
	database "github.com/jaysyanshar/godate-rest/internal/database"
	"github.com/jaysyanshar/godate-rest/repositories/account"
	"github.com/jaysyanshar/godate-rest/repositories/user"
	"github.com/jaysyanshar/godate-rest/routes"
	authSvc "github.com/jaysyanshar/godate-rest/services/auth"
)

func main() {
	cfg := config.Get()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	accountRepo := account.NewRepository(db)
	userRepo := user.NewRepository(db)
	authService := authSvc.NewService(cfg, accountRepo, userRepo)
	authController := authCtrl.NewController(authService)
	dashboardController := dashboardCtrl.NewController()

	router := routes.SetupRouter(authController, dashboardController)

	log.Printf("Starting server at port %s\n", cfg.AppPort)
	if err := http.ListenAndServe(":"+cfg.AppPort, router); err != nil {
		log.Fatal(err)
	}
}
