package responser

import "net/http"

type ErrorResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Error string `json:"error,omitempty"`
}

func GetBadRequest(msg string, err string) ErrorResponse {
	return ErrorResponse{
		Code:  http.StatusBadRequest,
		Msg:   msg,
		Error: err,
	}
}

func NotFound(msg string, err string) ErrorResponse {
	return ErrorResponse{
		Code:  http.StatusNotFound,
		Msg:   msg,
		Error: err,
	}
}
