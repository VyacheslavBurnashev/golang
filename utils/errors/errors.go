package errors

import "net/http"

type Errors struct {
	Message string
	Status  int
	Error   string
}

func ServerError(message string) *Errors {
	return &Errors{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func NewRequestError(message string) *Errors {
	return &Errors{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "internal_server_error",
	}
}
