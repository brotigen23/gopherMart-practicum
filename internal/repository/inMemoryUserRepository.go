package repository

import "github.com/brotigen23/gopherMart/internal/entity"

type InMemoryUserRepository struct {
}

func NewInMemoryUserRepository() UserRepository {
	return &InMemoryUserRepository{}
}

func (r *InMemoryUserRepository) GetUserByID() (*entity.User, error)    { return nil, nil }
func (r *InMemoryUserRepository) GetUserByLogin() (*entity.User, error) { return nil, nil }

func (r *InMemoryUserRepository) Save(*entity.User) error { return nil }
