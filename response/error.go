package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

func BadRequest(msg string) Response {
	return &ErrorResponse{
		Message: msg,
		Status:  http.StatusBadRequest,
	}
}

func Unauthorized(msg string) Response {
	return &ErrorResponse{
		Message: msg,
		Status:  http.StatusUnauthorized,
	}
}

func Forbidden(msg string) Response {
	return &ErrorResponse{
		Message: msg,
		Status:  http.StatusForbidden,
	}
}

func NotFound(msg string) Response {
	return &ErrorResponse{
		Message: msg,
		Status:  http.StatusNotFound,
	}
}

func MethodNotAllowed(msg string) Response {
	return &ErrorResponse{
		Message: msg,
		Status:  http.StatusMethodNotAllowed,
	}
}

func Conflict(msg string) Response {
	return &ErrorResponse{
		Message: msg,
		Status:  http.StatusConflict,
	}
}

func InternalServerError(msg string) Response {
	if msg == "" {
		msg = "Internal server error"
	}
	return &ErrorResponse{
		Message: msg,
		Status:  http.StatusInternalServerError,
	}
}

func ServiceUnavailable(msg string) Response {
	return &ErrorResponse{
		Message: msg,
		Status:  http.StatusServiceUnavailable,
	}
}

// Implementation of Response interface
func (e *ErrorResponse) Error() string {
	return e.Message
}

func (e *ErrorResponse) StatusCode() int {
	return e.Status
}

func (e *ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e *ErrorResponse) GetData() interface{} {
	return e.Data
}
