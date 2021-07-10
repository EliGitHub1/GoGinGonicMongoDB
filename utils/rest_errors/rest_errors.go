package rest_errors

import (
	"fmt"
	"net/http"
)

//Error will be from the following form
//if there will be a need to change in impl
//it deosn't going to do affect on function impl signutaruts
type RestErr interface {
	Message() string
	Status() int
	Description() string
}

type restErr struct {
	message     string `json:"message"`
	status_code int    `json:"status_code"`
	error_des   string `json:"error_des"`
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s ",
		e.message, e.status_code, e.error_des)
}

func (e restErr) Message() string {
	return e.message
}

func (e restErr) Status() int {
	return e.status_code
}

func (e restErr) Description() string {
	return e.error_des
}

func NewRestError(message string, status int, err string) RestErr {
	return restErr{
		message:     message,
		status_code: status,
		error_des:   err,
	}
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		message:     message,
		status_code: http.StatusBadRequest,
		error_des:   "bad_request",
	}
}
