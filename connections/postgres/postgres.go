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

type execer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
}

type ConnectPG struct {
	db execer
}

func ConnectToDB() (*ConnectPG, error) {
	db, err := sql.Open("postgres",
		"postgres://postgres:postgres@localhost:5432/complex_test_bd?sslmode=disable")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	err = db.Ping()
	pg := ConnectPG{db}
	if err != nil {
		pg.CloseConnectToBD()
		fmt.Print(err)
		return nil, err
	}
	return &pg, nil
}

func (pg *ConnectPG)CloseConnectToBD(){
	err := pg.db.Close()
	if err != nil {
		fmt.Print(err)
	}
}

func (pg *ConnectPG)GetAllBooks() ([]Book, error) {
	books := make([]Book, 0, 1)
	rows, err := pg.db.Query(`SELECT * FROM books`)
	if err != nil {
		return []Book{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.Id, &book.Name, &book.Author, &book.Year)
		if err != nil {
			return []Book{}, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return []Book{}, err
	}
	return books, nil
}

func (pg *ConnectPG)AddBook(name string, author string, year int16) (int32, error) {
	row := pg.db.QueryRow(`INSERT INTO books(name, author, year) 
                               VALUES ($1, $2, $3) RETURNING id`, name, author, year)
	var id int32
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pg *ConnectPG)DeleteBook(id int32) error {
	_, err := pg.db.Exec(`DELETE FROM books WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}