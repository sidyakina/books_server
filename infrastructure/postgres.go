package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sidyakina/books_server/adapters/postgres"
)

func ConnectToDB() (*postgres.ConnectDB, error) {
	db, err := sql.Open("postgres",
		"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	err = db.Ping()
	pg := postgres.ConnectDB{db}
	if err != nil {
		pg.CloseConnectToBD()
		fmt.Print(err)
		return nil, err
	}
	return &pg, nil
}
