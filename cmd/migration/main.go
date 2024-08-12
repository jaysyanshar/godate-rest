package main

import (
	"log"

	"github.com/jaysyanshar/godate-rest/config"
	db "github.com/jaysyanshar/godate-rest/internal/database"
	"github.com/jaysyanshar/godate-rest/models/dbmodel"
)

func main() {
	cfg := config.Get()
	dbClient, err := db.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	dbClient.AutoMigrate(&dbmodel.Account{}, &dbmodel.User{})
	defer dbClient.Close()
}
