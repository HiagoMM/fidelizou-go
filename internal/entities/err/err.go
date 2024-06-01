package err

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("Message: %s Status: %d", e.Message, e.Status)
}

func NewBadRequest(message string) *HttpError {
	return &HttpError{
		Status:  http.StatusBadRequest,
		Message: message,
	}
}

func NewUnauthorized(message string) *HttpError {
	return &HttpError{
		Status:  http.StatusUnauthorized,
		Message: message,
	}
}

func NewForbidden(message string) *HttpError {
	return &HttpError{
		Status:  http.StatusForbidden,
		Message: message,
	}
}

func NewNotFound(message string) *HttpError {
	return &HttpError{
		Status:  http.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerError(message string) *HttpError {
	return &HttpError{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}
