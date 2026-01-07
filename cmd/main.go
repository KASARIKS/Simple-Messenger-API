package main

import (
	"log"

	"github.com/kasariks/simple_messenger_api/cmd/api"
	"github.com/kasariks/simple_messenger_api/config"
)

func main() {
	server := api.NewServer(":"+config.Envs.Port, nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
