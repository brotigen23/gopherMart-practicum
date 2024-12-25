package service

import (
	"errors"
	"log"
	"time"

	"github.com/brotigen23/gopherMart/internal/dto"
	"github.com/brotigen23/gopherMart/internal/entity"
	"github.com/brotigen23/gopherMart/internal/repository"
)

type UserService struct {
	// TODO: БД
	// * Логгер для ошибок
	repository repository.Repository
}

func NewUserService(userRepo repository.Repository) *UserService {
	return &UserService{
		repository: userRepo,
	}
}

func (s *UserService) GetUserPasswordByLogin(login string) (string, error) {
	user, err := s.repository.GetUserByLogin(login)
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

func (s *UserService) IsUserExists(login string) bool {
	_, err := s.repository.GetUserByLogin(login)
	if err != nil && errors.Is(err, repository.ErrOrderNotFound) {
		return false
	}
	if err != nil {
		return false
	}
	return true
}

func (s *UserService) SaveUser(login string, password string) error {
	user := &entity.User{Login: login, Password: password}
	_, err := s.repository.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) SaveOrder(login string, orderNum string) error {
	// Получаем пользователя
	user, err := s.repository.GetUserByLogin(login)
	if err != nil {
		return err
	}
	// Проверяем наличие заказа в системе
	order, err := s.repository.GetOrderByNumber(orderNum)
	// Если заказа нет то продолжаем сохранять
	if err != nil && err.Error() != repository.ErrOrderNotFound.Error() {
		log.Println(err.Error())
		return err
	}
	// Если заказ есть
	if order != nil {
		// Проверям сохранен ли он тем пользователем
		if order.UserID != user.ID {
			return ErrOrderSavedByOtherUser
		}
		return ErrOrderAlreadySave
	}

	orderToSave := &entity.Order{
		UserID:     user.ID,
		Order:      orderNum,
		UploadedAt: time.Now(),
	}
	_, err = s.repository.SaveOrder(orderToSave)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetOrders(login string) ([]dto.Order, error) {
	ret := []dto.Order{}
	orders, err := s.repository.GetOrders(login)
	if err != nil {
		return nil, err
	}
	for _, order := range orders {
		item := dto.Order{
			Number:     order.Order,
			UploadedAt: order.UploadedAt,
		}
		ret = append(ret, item)
	}
	return ret, nil
}
