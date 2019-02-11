package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Book struct {
	Id     int32
	Name   string
	Author string
	Year   int16
}

type ConnectPG = sql.DB

func ConnectToDB() (*ConnectPG, error) {
	db, err := sql.Open("postgres",
		"postgres://postgres:postgres@localhost:5432/complex_test_bd?sslmode=disable")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		CloseConnectToBD(db)
		fmt.Print(err)
		return nil, err
	}
	return db, nil
}

func CloseConnectToBD(db *ConnectPG) {
	err := db.Close()
	if err != nil {
		fmt.Print(err)
	}
}

func GetAllBooks(db *ConnectPG) ([]Book, error) {
	count, err := countBooks(db)
	if err != nil {
		return []Book{}, err
	}
	books := make([]Book, count)
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Print(err)
		return []Book{}, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			fmt.Print(err)
		}
	}()
	i := 0
	for rows.Next() {
		err = rows.Scan(&books[i].Id, &books[i].Name, &books[i].Author, &books[i].Year)
		i ++
	}
	if rows.Err() != nil {
		return []Book{}, err
	}
	return books, nil
}

func countBooks(db *ConnectPG) (int32, error) {
	row := db.QueryRow("SELECT count(id) FROM books")
	var count int32
	err := row.Scan(&count)
	if err != nil {
		fmt.Printf("Error %v", err)
		return 0, err
	}
	return count, nil
}

func AddBook(db *ConnectPG, name string, author string, year int16) error {
	_, err := db.Exec("INSERT INTO books(\"name\", author, \"year\") VALUES ($1, $2, $3)", name, author, year)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(db *ConnectPG, id int32) error {
	_, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}