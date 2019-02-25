package use_case

import (
	"encoding/json"
	"github.com/sidyakina/books_server/domain"
	"time"
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

func (interactor *GetBookInteractor)GetAllBooks() []byte {
	books, err := interactor.bookRepo.GetAllBooks()
	result := make(map[string]interface{})
	result["books"] = books
	if err != nil {
		result["error"] = "Internal error"
	} else {
		result["error"] = ""
	}
	sResponse, _ := json.Marshal(result)
	return append(sResponse, byte('\n'))

}

func ErrorResult (msg string) []byte {
	result := make(map[string]interface{})
	result["error"] = msg
	sResponse, _ := json.Marshal(result)
	return append(sResponse, byte('\n'))
}

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

func (interactor *AddBookInteractor)AddBook (request []byte) []byte {
	prequest := domain.RequestAdd{}
	err := json.Unmarshal(request, &prequest)
	if err != nil {
		return ErrorResult("error while addBook cant't unmarshal Params")
	}
	name := prequest.Params.Name
	author := prequest.Params.Author
	year := prequest.Params.Year
	if int(year) > time.Now().Year() || year <= 0{
		return ErrorResult("error while addBook wrong Year")
	}
	if name == "" {
		return ErrorResult("error while addBook empty Name")
	}
	if author == "" {
		return ErrorResult("error while addBook empty Author")
	}
	id, err := interactor.bookRepo.AddBook(name, author, year)
	if err != nil {
		return ErrorResult ("Error while adding book!")
	}
	result := make(map[string]interface{})
	result["id"] = id
	result["error"] = ""
	sResponse, _ := json.Marshal(result)
	return append(sResponse, byte('\n'))
}

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

func (interactor *RemoveBookInteractor)RemoveBook (request [] byte) []byte {
	prequest := domain.RequestRemove{}
	err := json.Unmarshal(request, &prequest)
	if err != nil {
		return ErrorResult("error while deleteBook cant't unmarshal Params")
	}
	if prequest.Params.Id <= 0 {
		return ErrorResult("error while deleteBook wrong id")
	}
	err = interactor.bookRepo.DeleteBook(int32(prequest.Params.Id))
	if err != nil {
		return ErrorResult ("Error while deleting book!")
	}
	result := make(map[string]interface{})
	result["id"] = prequest.Params.Id
	result["error"] = ""
	sResponse, _ := json.Marshal(result)
	return append(sResponse, byte('\n'))
}