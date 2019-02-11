package main

import (
	"fmt"
	"github.com/sidyakina/books_server/connections/postgres"
)


func main() {
	pg, err := postgres.ConnectToDB()
	if err != nil {
		return
	}
	defer pg.CloseConnectToBD()
	err = pg.AddBook("new book 6", "new author 6 ", 1667)
	if err != nil {
		return
	}
	err = pg.DeleteBook(5)
	if err != nil {
		return
	}
	books, _ := pg.GetAllBooks()
	for i := 0; i < len(books); i++ {
		fmt.Printf("i = %v: id %v, %v, %v, %v \n", i + 1, books[i].Id, books[i].Name,
			books[i].Author, books[i].Year)
	}





}
