package main

import (
	"log"
	"net/http"

	"github.com/jaysyanshar/godate-rest/config"
	db "github.com/jaysyanshar/godate-rest/internal/database"
	"github.com/jaysyanshar/godate-rest/routes"
)

func main() {
	cfg := config.Get()
	_, err := db.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	router := routes.SetupRouter()

	log.Printf("Starting server at port %s\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}
