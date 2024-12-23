package repository

import "github.com/brotigen23/gopherMart/internal/entity"

type UserRepository interface {
	GetUserByID() (*entity.User, error)
	GetUserByLogin() (*entity.User, error)

	Save(*entity.User) error
}
