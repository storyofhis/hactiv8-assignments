package helper

import "net/http"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Payload any    `json:"payload,omitempty"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data"`
}

func BadRequestError(err error) *Response {
	return &Response{
		Status:  http.StatusBadRequest,
		Message: "BAD_REQUEST njir",
		Error:   err.Error(),
		Data:    nil,
	}
}

func InternalServerError(err error) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: "INTERNAL_SERVER_ERROR",
		Error:   err.Error(),
		Data:    nil,
	}
}

func DataNotFound(err error) *Response {
	return &Response{
		Status:  http.StatusNotFound,
		Message: "NO_DATA",
		Error:   err.Error(),
		Data:    nil,
	}
}

func DataConflict(err error) *Response {
	return &Response{
		Status:  http.StatusConflict,
		Message: "DUPLICATE_DATA",
		Error:   err.Error(),
		Data:    nil,
	}
}

func SuccessCreateResponse(payload interface{}, message string, data any) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: message,
		Payload: payload,
		Data:    data,
	}
}
func SuccessFindAllResponse(message string, data any) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		// Payload: payload,
		Data: data,
	}
}
