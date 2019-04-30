package main

import (
	"github.com/sidyakina/books_server/adapters/postgres"
	"github.com/sidyakina/books_server/adapters/server"
	"github.com/sidyakina/books_server/infrastructure"
	"github.com/sidyakina/books_server/use_cases"

	"log"
	"time"
)

const reconnects = 5

func main() {
	var pg *postgres.ConnectDB
	var err error
	for i := 0; i < reconnects; i++ {
		pg, err = infrastructure.ConnectToDB()
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

	sr, err := infrastructure.InitServer("localhost", "3333")
	if err != nil {
		log.Fatal(err)
	}
	defer sr.Stop()
	log.Println("Server started")
	sr.Start(handlers)
}
