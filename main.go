package main

import (
	"github.com/sidyakina/books_server/adapters/postgres"
	"github.com/sidyakina/books_server/adapters/server"
	"github.com/sidyakina/books_server/infrastructure"
	"github.com/sidyakina/books_server/use_cases"

	"log"
	"time"
)

func main() {
	config, err := setConfig()
	if err != nil {
		log.Fatal(err)
	}
	var pg *postgres.ConnectDB
	for i := 0; i < config.reconnect; i++ {
		pg, err = infrastructure.ConnectToDB(config.pgHost, config.pgPort, config.pgUser, config.pgPass, config.pgNameDB)
		if err == nil {
			break
		}
		log.Println("Err", err)
		time.Sleep(10 * time.Second)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer pg.CloseConnectToBD()
	getInt := use_cases.NewGetBookInteractor(pg)
	addInt := use_cases.NewAddBookInteractor(pg)
	remInt := use_cases.NewRemoveBookInteractor(pg)
	handlers := server.InitHandlers(getInt, addInt, remInt)

	sr, err := infrastructure.InitServer(config.serverPort)
	if err != nil {
		log.Fatal(err)
	}
	defer sr.Stop()
	log.Println("Server started")
	sr.Start(handlers)
}
