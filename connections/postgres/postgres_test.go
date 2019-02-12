package postgres

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"regexp"
	"testing"
)

func TestConnectPG_DeleteBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM books WHERE id = $1`)).
		WithArgs(5).WillReturnResult(sqlmock.NewResult(1, 1))

	pg := ConnectPG{db}
	if err = pg.DeleteBook(5); err != nil {
		t.Errorf("error was not expected while deleting book: %s", err)
	}
}

func TestConnectPG_DeleteBook2(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM books WHERE id = $1`)).
		WithArgs(5).WillReturnError(sqlmock.ErrCancelled)

	pg := ConnectPG{db}
	if err = pg.DeleteBook(5); err != sqlmock.ErrCancelled {
		t.Errorf("error was expected while deleting book: %s", err)
	}
}

func TestConnectPG_AddBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO books(name, author, year) VALUES ($1, $2, $3)`)).
		WithArgs("newBook", "author33", 1333).WillReturnResult(sqlmock.NewResult(1, 1))

	pg := ConnectPG{db}
	if err = pg.AddBook("newBook", "author33", 1333); err != nil {
		t.Errorf("error was not expected while adding book: %s", err)
	}
}

func TestConnectPG_AddBook2(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO books(name, author, year) VALUES ($1, $2, $3)`)).
		WithArgs("newBook", "author33", 1333).WillReturnError(sqlmock.ErrCancelled)

	pg := ConnectPG{db}
	if err = pg.AddBook("newBook", "author33", 1333); err != sqlmock.ErrCancelled {
		t.Errorf("error was expected while adding book: %s", err)
	}
}

func TestConnectPG_GetAllBooks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM books`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "author", "year"}).
			AddRow(1, "newBook", "newAuthor", 1333))

	pg := ConnectPG{db}
	books, err := pg.GetAllBooks()
	if err != nil {
		t.Errorf("error was not expected while getting books: %s", err)
	}
	expected := []Book{{1, "newBook", "newAuthor", 1333}}
	assert.Equal(t, expected, books)
}

func TestConnectPG_GetAllBooks2(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM books`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "author", "year"}).
			AddRow(1, "newBook", "newAuthor", 1333).
			AddRow(2, "newBook2", "newAuthor2", 1222).
			AddRow(3, "newBook3", "newAuthor3", 1555))

	pg := ConnectPG{db}
	books, err := pg.GetAllBooks()
	if err != nil {
		t.Errorf("error was not expected while getting books: %s", err)
	}
	expected := []Book{{1, "newBook", "newAuthor", 1333},
		{2, "newBook2", "newAuthor2", 1222},
		{3, "newBook3", "newAuthor3", 1555}}
	assert.Equal(t, expected, books)
}

func TestConnectPG_GetAllBooks3(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM books`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "author", "year"}).
			AddRow(1, "newBook", "newAuthor", 1333).
			AddRow(2, "newBook2", "newAuthor2", 1222).RowError(1, sqlmock.ErrCancelled))

	pg := ConnectPG{db}
	books, err := pg.GetAllBooks()
	if err != sqlmock.ErrCancelled {
		t.Errorf("error was expected while getting books: %s", err)
	}
	assert.Equal(t, []Book{}, books)
}

func TestConnectPG_GetAllBooks4(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM books`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "author"}).
			AddRow(1, "newBook", "newAuthor"))

	pg := ConnectPG{db}
	books, err := pg.GetAllBooks()
	assert.Error(t, err)
	assert.Equal(t, []Book{}, books)
}


func TestConnectPG_GetAllBooks5(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM books`)).
		WillReturnError(sqlmock.ErrCancelled)

	pg := ConnectPG{db}
	books, err := pg.GetAllBooks()
	if err != sqlmock.ErrCancelled {
		t.Errorf("error was expected while getting books: %s", err)
	}
	assert.Equal(t, []Book{}, books)
}

func TestConnectPG_GetAllBooks6(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM books`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "author"}))

	pg := ConnectPG{db}
	books, err := pg.GetAllBooks()
	assert.NoError(t, err)
	assert.EqualValues(t, []Book{}, books)
}