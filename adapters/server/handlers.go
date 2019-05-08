package server

import (
	"encoding/json"
	"fmt"
	"github.com/sidyakina/books_server/domain"
)

type handlerGet interface {
	GetAllBooks() ([]domain.Book, string)
}

type handlerAdd interface {
	AddBook(request domain.RequestAdd) (int32, string)
}

type handlerRemove interface {
	RemoveBook(request domain.RequestRemove) (int32, string)
}

type Handlers struct {
	get    handlerGet
	add    handlerAdd
	remove handlerRemove
}

func InitHandlers(hget handlerGet, hadd handlerAdd, hremove handlerRemove) *Handlers {
	return &Handlers{get: hget, add: hadd, remove: hremove}
}

func handleRequest(buf []byte, handlers *Handlers) []byte {
	request := make(map[string]interface{})
	result := make(map[string]interface{})
	err := json.Unmarshal(buf, &request)
	if err != nil {
		fmt.Println(err)
		result["error"] = "error: can't unmarshal request."
		response, _ := json.Marshal(result)
		return response
	}
	fmt.Printf("get request %v\n", request)
	cmd, _ := request["cmd"]
	result["cmd"] = cmd
	switch cmd {
	case "getAllBooks":
		books, err := handlers.get.GetAllBooks()
		result["books"] = books
		result["error"] = err
	case "addBook":
		request := domain.RequestAdd{}
		err := json.Unmarshal(buf, &request)
		if err != nil {
			fmt.Print(err)
			result["id"] = 0
			result["error"] = "error: can't unmarshal request."
		} else {
			id, err := handlers.add.AddBook(request)
			result["id"] = id
			result["error"] = err
		}
	case "deleteBook":
		request := domain.RequestRemove{}
		err := json.Unmarshal(buf, &request)
		if err != nil {
			fmt.Println(err)
			result["error"] = "error while deleteBook: cant't unmarshal Params"
		} else {
			id, err := handlers.remove.RemoveBook(request)
			result["id"] = id
			result["error"] = err
		}
	default:
		result["error"] = fmt.Sprintf("cmd: %v is not defined", cmd)
	}
	//fmt.Printf("send response %v\n", response)
	response, _ := json.Marshal(result)
	return response
}
