package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func OK(msg string, data interface{}) Response {
	return &SuccessResponse{
		Message: msg,
		Status:  http.StatusOK,
		Data:    data,
	}
}

func Created(msg string, data interface{}) Response {
	return &SuccessResponse{
		Message: msg,
		Status:  http.StatusCreated,
		Data:    data,
	}
}

func Accepted(msg string, data interface{}) Response {
	return &SuccessResponse{
		Message: msg,
		Status:  http.StatusAccepted,
		Data:    data,
	}
}

func NonAuthoritativeInfo(msg string, data interface{}) Response {
	return &SuccessResponse{
		Message: msg,
		Status:  http.StatusNonAuthoritativeInfo,
		Data:    data,
	}
}

func NoContent(msg string, data interface{}) Response {
	return &SuccessResponse{
		Message: msg,
		Status:  http.StatusNoContent,
		Data:    data,
	}
}

func ResetContent(msg string, data interface{}) Response {
	return &SuccessResponse{
		Message: msg,
		Status:  http.StatusResetContent,
		Data:    data,
	}
}

func (s *SuccessResponse) Error() string {
	return ""
}

func (s *SuccessResponse) StatusCode() int {
	return s.Status
}

func (s *SuccessResponse) GetBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SuccessResponse) GetData() interface{} {
	return s.Data
}
