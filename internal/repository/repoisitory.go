package repository

import "github.com/brotigen23/gopherMart/internal/entity"

type Repository interface {
	GetUserByID(id int) (*entity.User, error)
	GetUserByLogin(login string) (*entity.User, error)

	SaveUser(user *entity.User) (*entity.User, error)

	GetOrders(login string) ([]entity.Order, error)
	SaveOrder(*entity.Order) (*entity.Order, error)
}
