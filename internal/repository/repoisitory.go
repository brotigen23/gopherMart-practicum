package repository

import "github.com/brotigen23/gopherMart/internal/entity"

type Repository interface {
	// Users
	GetUserByID(id int) (*entity.User, error)
	GetUserByLogin(login string) (*entity.User, error)

	SaveUser(user *entity.User) (*entity.User, error)
	UpdateUserBalance(sum float32) error

	//Orders
	GetOrders(login string) ([]entity.Order, error)
	GetOrderByNumber(orderNum string) (*entity.Order, error)
	UpdateOrderStatus(status string, order string) error

	SaveOrder(*entity.Order) (*entity.Order, error)

	// Withdrawals
	GetUserWithdrawals(user *entity.User) ([]entity.Withdraw, error)
}
