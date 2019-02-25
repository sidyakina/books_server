package use_cases

import (
	"fmt"
	"github.com/sidyakina/books_server/domain"
	"time"
)

type BookRepoAdd interface {
	AddBook(name string, author string, year int16) (int32, error)
}

type AddBookInteractor struct {
	bookRepo BookRepoAdd
}

func NewAddBookInteractor(bookRepo BookRepoAdd) *AddBookInteractor {
	return &AddBookInteractor{
		bookRepo: bookRepo,
	}
}

func (interactor *AddBookInteractor)AddBook (request domain.RequestAdd) (int32, string) {
	name := request.Params.Name
	author := request.Params.Author
	year := request.Params.Year
	if int(year) > time.Now().Year() || year <= 0{
		return 0, "error while addBook: wrong year"
	}
	if name == "" {
		return 0, "error while addBook: empty name"
	}
	if author == "" {
		return 0, "error while addBook: empty author"
	}
	id, err := interactor.bookRepo.AddBook(name, author, year)
	if err != nil {
		fmt.Print(err)
		return id, "internal error in db while addBook"
	}
	return id, ""
}

