package handler

import "errors"

var (
	ErrInternalServer = errors.New("internal server error")
	ErrBadRequest     = errors.New("bad request")
	ErrUserExists = errors.New("user exists")
)
