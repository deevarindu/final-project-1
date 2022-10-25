package views

import (
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BadRequestResponse(err error) *Response {
	return &Response{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	}
}

func InternalServerError(err error) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
	}
}

func DataConflictResponse(err error) *Response {
	return &Response{
		Status:  http.StatusConflict,
		Message: err.Error(),
	}
}

func SuccessCreateResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: message,
		Data:    payload,
	}
}

func SuccessResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    payload,
	}
}

func SuccessDeleteResponse(message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
	}
}
