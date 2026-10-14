package errapp

import (
	"errors"
	"net/http"
)

type ErrApp struct {
	ErrEmailAlreadyInUse error
	ErrUserNotFound      error
}

type ErrAppResult struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Err     error  `json:"error"`
}

func NewErrApp() *ErrApp {
	return &ErrApp{
		ErrEmailAlreadyInUse: errors.New("E-mail em uso"),
		ErrUserNotFound:      errors.New("Usuário não encontrado"),
	}
}

func (e *ErrApp) ErrResponse(msg string, code int, err error) errAppResult {
	return errAppResult{Message: msg, Code: code, Err: err}
}

func (e *ErrApp) BadRequestError(msg string, err error) errAppResult {
	return errAppResult{Message: msg, Err: err, Code: http.StatusBadRequest}
}

func (e *ErrApp) InternalServerError(msg string, err error) errAppResult {
	return errAppResult{Message: msg, Err: err, Code: http.StatusInternalServerError}
}

func (e *ErrApp) NotFoundError(msg string, err error) errAppResult {
	return errAppResult{Message: msg, Err: err, Code: http.StatusNotFound}
}

func (e *ErrApp) ForbiddenrError(msg string, err error) errAppResult {
	return errAppResult{Message: msg, Err: err, Code: http.StatusForbidden}
}

func (e *ErrApp) UnauthorizedError(msg string, err error) errAppResult {
	return errAppResult{Message: msg, Err: err, Code: http.StatusUnauthorized}
}
