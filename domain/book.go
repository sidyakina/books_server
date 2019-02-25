package domain

type Book struct {
	Id     int32  `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int16  `json:"year"`
}

