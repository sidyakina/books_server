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
	return sResponse

}

func ErrorResult (msg string) []byte {
	result := make(map[string]interface{})
	result["error"] = msg
	sResponse, _ := json.Marshal(result)
	return sResponse
}