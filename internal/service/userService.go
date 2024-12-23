package service

import "github.com/brotigen23/gopherMart/internal/repository"

type UserService struct {
	// TODO: БД
	// * Логгер для ошибок
	userRepository repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repository.NewInMemoryUserRepository(),
	}
}

func (s *UserService) IsUserExists(login string) bool {
	s.userRepository.GetUserByLogin()
	return false
}

func (s *UserService) Save(login string, password string) error {
	// * Возможно стоит захешировать пароль
	return nil
}
