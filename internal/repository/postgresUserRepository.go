package repository

import (
	"log"

	"github.com/brotigen23/gopherMart/internal/entity"
)

func (r *postgresRepository) GetUserByID(id int) (*entity.User, error) {
	return nil, nil
}

func (r *postgresRepository) GetUserByLogin(login string) (*entity.User, error) {
	query := r.db.QueryRow(`SELECT id, login, password, balance FROM Users WHERE login = $1`, login)
	var ID int
	var Login string
	var Password string
	var Balance float32
	err := query.Scan(&ID, &Login, &Password, &Balance)
	if err != nil && err == ErrUserNotFound {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:       ID,
		Login:    Login,
		Password: Password,
		Balance:  Balance,
	}, nil
}

func (r *postgresRepository) SaveUser(user *entity.User) (*entity.User, error) {
	query := "INSERT INTO Users(login, password) VALUES($1, $2) RETURNING ID"
	var (
		id int
	)
	err := r.db.QueryRow(query, user.Login, user.Password).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:       id,
		Login:    user.Login,
		Password: user.Password,
		//Balance:  0,
	}, nil
}

func (r *postgresRepository) UpdateUserBalance(user *entity.User, sum float32) error {
	query := "UPDATE Users SET balance = balance + $1 WHERE id = $2"

	_, err := r.db.Exec(query, sum, user.ID)

	if err != nil {
		return err
	}
	log.Println(r.GetUserByLogin(user.Login))
	return nil
}
