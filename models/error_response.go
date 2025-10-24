package models

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
	Filed        string `json:"filed"`
}
