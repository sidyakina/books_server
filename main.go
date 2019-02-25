package main

import (
	"github.com/sidyakina/books_server/adaptors/server"
	"github.com/sidyakina/books_server/infrastructure"
	"github.com/sidyakina/books_server/use_cases"

	"log"
)

func main() {
	pg, err := infrastructure.ConnectToDB()
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
	sr.Start(handlers)
}
