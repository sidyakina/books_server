package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sidyakina/books_server/adapters/postgres"
)

func ConnectToDB(host, port, user, pass, dbname string) (*postgres.ConnectDB, error) {
	//db, err := sql.Open("postgres",
	//	"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	// "postgres://user:pass@localhost/dbname"
	db, err := sql.Open("postgres",
		"postgres://" + user + ":" + pass + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable")
	if err != nil {
		fmt.Println("err: ", err)
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
