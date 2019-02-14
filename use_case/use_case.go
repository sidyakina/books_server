package use_case

import (
	"encoding/json"
	"github.com/sidyakina/books_server/connections/postgres"
	"time"
)

type requestAdd struct {
	Params paramsAdd `json:"params"`
}

type paramsAdd struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int16  `json:"year"`
}

type requestRemove struct {
	Params paramsRemove `json:"params"`
}

type paramsRemove struct {
	Id int32 `json:"id"`
}

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

func AddBook (pg *postgres.ConnectPG, request []byte) []byte {
	prequest := requestAdd{}
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
	id, err := pg.AddBook(name, author, year)
	if err != nil {
		return ErrorResult ("Error while adding book!")
	}
	result := make(map[string]interface{})
	result["id"] = id
	result["error"] = ""
	sResponse, _ := json.Marshal(result)
	return append(sResponse, byte('\n'))
}

func DeleteBook (pg *postgres.ConnectPG, request [] byte) []byte {
	prequest := requestRemove{}
	err := json.Unmarshal(request, &prequest)
	if err != nil {
		return ErrorResult("error while deleteBook cant't unmarshal Params")
	}
	if prequest.Params.Id <= 0 {
		return ErrorResult("error while deleteBook wrong id")
	}
	err = pg.DeleteBook(int32(prequest.Params.Id))
	if err != nil {
		return ErrorResult ("Error while deleting book!")
	}
	result := make(map[string]interface{})
	result["id"] = prequest.Params.Id
	result["error"] = ""
	sResponse, _ := json.Marshal(result)
	return append(sResponse, byte('\n'))
}