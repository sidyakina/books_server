package use_cases

import (
	"fmt"

	"github.com/sidyakina/books_server/domain"
)

type BookRepoGet interface {
	GetAllBooks() ([]domain.Book, error)
}

type GetBookInteractor struct {
	bookRepo BookRepoGet
}

func NewGetBookInteractor(bookRepo BookRepoGet) *GetBookInteractor {
	return &GetBookInteractor{
		bookRepo: bookRepo,
	}
}

func (interactor *GetBookInteractor) GetAllBooks() ([]domain.Book, string) {
	books, err := interactor.bookRepo.GetAllBooks()
	if err != nil {
		fmt.Print(err)
		return books, "internal error in db while getBook"
	}
	return books, ""
}
