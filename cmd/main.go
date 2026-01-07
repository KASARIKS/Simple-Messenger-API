package main

import (
	"log"

	"github.com/kasariks/simple_messenger_api/cmd/api"
	"github.com/kasariks/simple_messenger_api/config"
	"github.com/kasariks/simple_messenger_api/db"
)

func main() {
	db, err := db.NewSQLStorage(config.Envs.DBName)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(":"+config.Envs.Port, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
