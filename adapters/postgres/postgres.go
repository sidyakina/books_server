package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sidyakina/books_server/domain"
)

type execer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
}

type ConnectDB struct {
	DB execer
}

func (pg *ConnectDB) CloseConnectToBD() {
	err := pg.DB.Close()
	if err != nil {
		fmt.Print(err)
	}
}

func (pg *ConnectDB) GetAllBooks() ([]domain.Book, error) {
	books := make([]domain.Book, 0, 1)
	rows, err := pg.DB.Query(`SELECT * FROM books`)
	if err != nil {
		fmt.Println(err)
		return []domain.Book{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var book domain.Book
		err = rows.Scan(&book.ID, &book.Name, &book.Author, &book.Year)
		if err != nil {
			return []domain.Book{}, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return []domain.Book{}, err
	}
	return books, nil
}

func (pg *ConnectDB) AddBook(name string, author string, year int16) (int32, error) {
	row := pg.DB.QueryRow(`INSERT INTO books(name, author, year) 
                               VALUES ($1, $2, $3) RETURNING id`, name, author, year)
	var id int32
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pg *ConnectDB) DeleteBook(id int32) error {
	_, err := pg.DB.Exec(`DELETE FROM books WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
