package handler

import "errors"

var (
	ErrInternalServer = errors.New("internal server error")
	ErrBadRequest     = errors.New("bad request")
	ErrUserExists     = errors.New("user exists")
	ErrJWT            = errors.New("auth error")
	ErrWrongLogin     = errors.New("login incorrect")
	ErrWrongPassword  = errors.New("passwod incorrect")
)
