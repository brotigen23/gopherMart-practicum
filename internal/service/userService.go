package service

import (
	"github.com/brotigen23/gopherMart/internal/entity"
	"github.com/brotigen23/gopherMart/internal/repository"
)

type UserService struct {
	// TODO: БД
	// * Логгер для ошибок
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

func (s *UserService) IsUserExists(login string) bool {
	_, err := s.userRepository.GetUserByLogin(login)
	if err != nil && err == repository.ErrUserNotFound {
		return false
	}
	if err != nil {
		return false
	}
	return true
}

func (s *UserService) Save(login string, password string) error {
	user := &entity.User{Login: login, Password: password}
	_, err := s.userRepository.Save(user)
	if err != nil {
		return err
	}
	return nil
}
