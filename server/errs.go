package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error is service error.
type Error struct {
	Code    int    `json:"code"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

// Error returns a text message corresponding to the given error.
func (e *Error) Error() string {
	return e.Message
}

// StatusCode returns an HTTP status code corresponding to the given erro.
func (e *Error) StatusCode() int {
	return e.Code
}

// MarshalJSON marshal error into valid JSON.
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(&e)
}

func errBadRequest(format string, v ...interface{}) error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf(format, v...),
		Level:   "api",
	}
}

func errNotFound(format string, v ...interface{}) error {
	return &Error{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf(format, v...),
		Level:   "api",
	}
}

func errForbidden(format string, v ...interface{}) error {
	return &Error{
		Code:    http.StatusForbidden,
		Message: fmt.Sprintf(format, v...),
		Level:   "api",
	}
}
