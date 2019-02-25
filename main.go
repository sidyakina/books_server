package main

import (
	"github.com/sidyakina/books_server/adaptor/server"
	"github.com/sidyakina/books_server/infrastructure"
	"github.com/sidyakina/books_server/use_case"

	"log"
)

func main() {
	pg, err := infrastructure.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer pg.CloseConnectToBD()
	getInt := use_case.NewGetBookInteractor(pg)
	addInt := use_case.NewAddBookInteractor(pg)
	remInt := use_case.NewRemoveBookInteractor(pg)
	handlers := server.InitHandlers(getInt, addInt, remInt)

	sr, err := infrastructure.InitServer("localhost", "3333")
	if err != nil {
		log.Fatal(err)
	}
	defer sr.Stop()
	sr.Start(handlers)
}
