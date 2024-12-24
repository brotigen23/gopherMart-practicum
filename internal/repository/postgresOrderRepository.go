package repository

import "github.com/brotigen23/gopherMart/internal/entity"

type postgresOrderRepository struct {
}

func (r *postgresOrderRepository) GetOrderByID() (*entity.Order, error) {
	return nil, nil
}
func (r *postgresOrderRepository) Save(*entity.Order) (*entity.Order, error) {
	return nil, nil
}
