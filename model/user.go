package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ErrorReturn struct {
	ErrorMessage string `json:"error_message"`
}
