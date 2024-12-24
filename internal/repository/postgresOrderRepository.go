package repository

import (
	"time"

	"github.com/brotigen23/gopherMart/internal/entity"
)

func (r *postgresRepository) GetOrders(login string) ([]entity.Order, error) {
	ret := []entity.Order{}

	user, err := r.GetUserByLogin(login)
	if err != nil {
		return nil, err
	}
	query, err := r.db.Query(`SELECT * FROM Orders WHERE user_id = $1 ORDER BY uploaded_at `, user.ID)
	if err != nil {
		return nil, err
	}

	var ID int
	var UserID int
	var Order string
	var UploadedAt time.Time
	for query.Next() {
		err := query.Scan(&ID, &UserID, &Order, &UploadedAt)
		if err != nil {
			return nil, err
		}
		ret = append(ret, entity.Order{
			ID:         ID,
			UserID:     UserID,
			Order:      Order,
			UploadedAt: UploadedAt,
		})
	}
	return ret, nil
}
func (r *postgresRepository) SaveOrder(order *entity.Order) (*entity.Order, error) {
	query := `INSERT INTO Orders(user_id, "order", uploaded_at) VALUES($1, $2, $3) RETURNING ID`
	time := time.Now()
	var (
		id int
	)
	err := r.db.QueryRow(query, order.UserID, order.Order, time).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &entity.Order{
		ID:         id,
		UserID:     order.UserID,
		Order:      order.Order,
		UploadedAt: time,
	}, nil
}
