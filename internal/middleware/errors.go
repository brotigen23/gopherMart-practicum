package middleware

import "errors"

var (
	ErrNotValidJSON = errors.New("JSON is not valid")
	ErrContentType  = errors.New("not JSON")
)
