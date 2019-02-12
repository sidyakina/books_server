package main

import (
	"github.com/sidyakina/books_server/connections/postgres"
	"github.com/sidyakina/books_server/connections/server"
	"log"
)

func main() {
	pg, err := postgres.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer pg.CloseConnectToBD()
	sr, err := server.Init(pg, "localhost", "3333")
	if err != nil {
		log.Fatal(err)
	}
	defer sr.Stop()
	sr.Start()
}
