package main

import (
	"log"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/db"
	"scheduling-app-back-end/internal/server"
	"scheduling-app-back-end/internal/utils"
)

func main() {

	config, err := utils.LoadConfig(".")
	db.ConnectToPostgres()
	mailChan := make(chan models.MailData)
	if err != nil {
		log.Fatal(err)
	}
	newServer, err := server.NewServer(config, mailChan)

	go func() {
		msg := <-newServer.MailChan
		sendMsg(msg)
	}()

	if err != nil {
		log.Fatal(err)
	}
	err = newServer.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer close(newServer.MailChan)
}
