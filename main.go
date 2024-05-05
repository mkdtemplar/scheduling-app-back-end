package main

import (
	"log"
	"scheduling-app-back-end/internal/repository/db"
	"scheduling-app-back-end/internal/server"
	"scheduling-app-back-end/internal/utils"
)

func main() {

	config, err := utils.LoadConfig(".")
	db.ConnectToPostgres()
	if err != nil {
		log.Fatal(err)
	}
	newServer, err := server.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}
	err = newServer.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
