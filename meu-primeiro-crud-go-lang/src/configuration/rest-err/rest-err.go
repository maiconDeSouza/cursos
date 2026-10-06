package resterr

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field    string `json:"field"`
	Messsage string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(msg, err string, code int, causes []Causes) *RestErr {
	return &RestErr{Message: msg, Err: err, Code: code, Causes: causes}
}

func NewBadRequestError(msg string) *RestErr {
	return &RestErr{Message: msg, Err: "Bad Request", Code: http.StatusBadRequest}
}

func NewBadRequestValidationError(msg string, causes []Causes) *RestErr {
	return &RestErr{Message: msg, Err: "Bad Request", Code: http.StatusBadRequest, Causes: causes}
}

func NewInternalServerError(msg string) *RestErr {
	return &RestErr{Message: msg, Err: "Internal Server", Code: http.StatusInternalServerError}
}

func NewNotFoundError(msg string) *RestErr {
	return &RestErr{Message: msg, Err: "Not Found", Code: http.StatusNotFound}
}

func NewForbiddenrError(msg string) *RestErr {
	return &RestErr{Message: msg, Err: "Forbidden", Code: http.StatusForbidden}
}
