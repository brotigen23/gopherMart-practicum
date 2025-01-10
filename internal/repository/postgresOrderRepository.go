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
	query := `	SELECT id, user_id, "order", uploaded_at 
				FROM Orders 
				WHERE user_id = $1 
				ORDER BY uploaded_at `
	rows, err := r.db.Query(query, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ID int
	var UserID int
	var Order string
	var UploadedAt time.Time
	for rows.Next() {
		err := rows.Scan(&ID, &UserID, &Order, &UploadedAt)
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
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return ret, nil
}
func (r *postgresRepository) GetOrderByNumber(orderNum string) (*entity.Order, error) {
	query := `	SELECT id, user_id, "order", uploaded_at 
				FROM Orders 
				WHERE "order" = $1`
	row := r.db.QueryRow(query, orderNum)
	var ID int
	var UserID int
	var OrderNum string
	var UploadedAt time.Time
	err := row.Scan(&ID, &UserID, &OrderNum, &UploadedAt)
	if err != nil {
		return nil, err
	}
	return &entity.Order{
		ID:         ID,
		UserID:     UserID,
		Order:      OrderNum,
		UploadedAt: UploadedAt,
	}, nil
}

func (r *postgresRepository) SaveOrder(order *entity.Order) (*entity.Order, error) {
	query := `	INSERT INTO Orders(user_id, "order", uploaded_at) 
				VALUES($1, $2, $3) 
				RETURNING ID`
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
