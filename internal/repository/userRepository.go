package repository

import "github.com/brotigen23/gopherMart/internal/entity"

type UserRepository interface {
	GetUserByID(id int) (*entity.User, error)
	GetUserByLogin(login string) (*entity.User, error)
	
	Save(user *entity.User) (*entity.User, error)
}
