package repository

import (
	"time"

	"github.com/brotigen23/gopherMart/internal/entity"
)

func (r *postgresRepository) GetUserWithdrawals(user *entity.User) ([]entity.Withdraw, error) {
	ret := []entity.Withdraw{}
	query := `SELECT id, user_id, "order", sum, processed_at FROM withdrawals WHERE user_id = $1 ORDER BY processed_at`
	rows, err := r.db.Query(query, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ID int
	var UserID int
	var order string
	var Sum float32
	var ProccessedAt time.Time
	for rows.Next() {
		err = rows.Scan(&ID, &UserID, order, &Sum, &ProccessedAt)
		if err != nil {
			return nil, err
		}
		ret = append(ret, entity.Withdraw{
			ID:           ID,
			UserID:       UserID,
			Order:        order,
			Sum:          Sum,
			ProccessedAt: ProccessedAt,
		})
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return ret, nil
}

func (r *postgresRepository) SaveWithdraw(user *entity.User, withdraw *entity.Withdraw) error {
	query := `INSERT INTO Withdrawals(user_id, "order", sum, processed_at) VALUES($1, $2, $3, $4)`
	_, err := r.db.Exec(query, user.ID, withdraw.Order, withdraw.Sum, withdraw.ProccessedAt)
	if err != nil {
		return err
	}
	return nil
}
