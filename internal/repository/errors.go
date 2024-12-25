package repository

import "errors"

// * Users
var (
	ErrUserNotFound = errors.New("sql: no rows in result set")
	ErrUserExists   = errors.New("user already exists")
)

// * Orders
var (
	ErrOrderNotFound = errors.New("sql: no rows in result set")
	ErrOrderNotValid = errors.New("order not valid")
)
