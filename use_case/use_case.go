package use_case

import (
	"encoding/json"
	"github.com/sidyakina/books_server/connections/postgres"
)

func GetAllBooks(pg *postgres.ConnectPG) []byte {
	books, err := pg.GetAllBooks()
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

func AddBook (pg *postgres.ConnectPG, request map[string]interface{}) []byte {
	p, exist := request["params"]
	if !exist {
		return ErrorResult("cmd: addBook must have params")
	}
	params, ok := p.(map[string]interface{})
	if !ok {
		return ErrorResult("cmd: addBook must have params")
	}
	name, exist1 :=  params["name"]
	author, exist2 :=  params["author"]
	year, exist3 := params["year"]
	if !(exist1 && exist2 && exist3) {
		return ErrorResult("cmd: addBook must have name, author, year")
	}
	n, ok1 := name.(string)
	a, ok2 := author.(string)
	y, ok3 := year.(float64)
	if !(ok1 && ok2 && ok3) {
		return ErrorResult("cmd: addBook, wrong type parameters")
	}
	id, err := pg.AddBook(n, a, int16(y))
	if err != nil {
		return ErrorResult ("Error while adding book!")
	}
	result := make(map[string]interface{})
	result["id"] = id
	result["error"] = ""
	sResponse, _ := json.Marshal(result)
	return append(sResponse, byte('\n'))
}

func DeleteBook (pg *postgres.ConnectPG, request map[string]interface{}) []byte {
	p, exist := request["params"]
	if !exist {
		return ErrorResult("cmd: deleteBook must have params")
	}
	params, ok := p.(map[string]interface{})
	if !ok {
		return ErrorResult("cmd: deleteBook must have params")
	}
	id, exist :=  params["id"]
	if !exist {
		return ErrorResult("cmd: deleteBook must have id")
	}
	i, ok := id.(float64)
	if !ok {
		return ErrorResult("cmd: deleteBook, wrong type parameter")
	}
	err := pg.DeleteBook(int32(i))
	if err != nil {
		return ErrorResult ("Error while deleting book!")
	}
	result := make(map[string]interface{})
	result["id"] = id
	result["error"] = ""
	sResponse, _ := json.Marshal(result)
	return append(sResponse, byte('\n'))
}