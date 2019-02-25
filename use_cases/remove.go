package use_cases

import (
	"fmt"
	"github.com/sidyakina/books_server/domain"
)

type BookRepoRemove interface {
	DeleteBook(id int32) error
}

type RemoveBookInteractor struct {
	bookRepo BookRepoRemove
}

func NewRemoveBookInteractor(bookRepo BookRepoRemove) *RemoveBookInteractor {
	return &RemoveBookInteractor{
		bookRepo: bookRepo,
	}
}

func (interactor *RemoveBookInteractor)RemoveBook (request domain.RequestRemove) (int32, string) {
	id := request.Params.Id
	if id <= 0 {
		return 0, "error while deleteBook wrong id"
	}
	err := interactor.bookRepo.DeleteBook(int32(id))
	if err != nil {
		fmt.Println(err)
		return 0, "Internal error in db while deleting book!"
	}
	return id, ""
}