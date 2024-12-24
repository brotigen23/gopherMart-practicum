package service

import (
	"time"

	"github.com/brotigen23/gopherMart/internal/entity"
	"github.com/brotigen23/gopherMart/internal/repository"
)

type UserService struct {
	// TODO: БД
	// * Логгер для ошибок
	userRepository     repository.UserRepository
	orderRepository    repository.OrderRepository
	withdrawRepository repository.WithdrawRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

func (s *UserService) GetUserPasswordByLogin(login string) (string, error) {
	user, err := s.userRepository.GetUserByLogin(login)
	if err != nil {
		return "", err
	}
	return user.Password, nil
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

func (s *UserService) SaveUser(login string, password string) error {
	user := &entity.User{Login: login, Password: password}
	_, err := s.userRepository.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) SaveOrder(login string, orderNum string) error {
	user, err := s.userRepository.GetUserByLogin(login)
	if err != nil {
		return err
	}
	order := &entity.Order{
		UserID:     user.ID,
		Order:      orderNum,
		UploadedAt: time.Now(),
	}
	_, err = s.orderRepository.Save(order)
	if err != nil {
		return err
	}
	return nil
}
