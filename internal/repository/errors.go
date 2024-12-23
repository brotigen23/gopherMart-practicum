package repository

import "errors"

// * Users
var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
)

// * Orders
var (
	ErrOrderNotFound = errors.New("order not found")
	ErrOrderNotValid = errors.New("order not valid")
)
