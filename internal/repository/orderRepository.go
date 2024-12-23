package repository

import "github.com/brotigen23/gopherMart/internal/entity"

type OrderRepository interface {
	GetOrderByID() (*entity.Order, error)
}
