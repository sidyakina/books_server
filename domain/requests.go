package domain

type RequestAdd struct {
	Params ParamsAdd `json:"params"`
}

type ParamsAdd struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int16  `json:"year"`
}

type RequestRemove struct {
	Params ParamsRemove `json:"params"`
}

type ParamsRemove struct {
	ID int32 `json:"id"`
}
