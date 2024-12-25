package repository

import (
	"time"

	"github.com/brotigen23/gopherMart/internal/entity"
)

func (r *postgresRepository) GetUserWithdrawals(user *entity.User) ([]entity.Withdraw, error) {
	ret := []entity.Withdraw{}
	query := `SELECT id, user_id, sum, proccessed_at FROM withdrawals WHERE user_id = $1 ORDER BY proccessed_at`
	q, err := r.db.Query(query, user.ID)
	if err != nil {
		return nil, err
	}

	var ID int
	var UserID int
	var Sum float32
	var ProccessedAt time.Time
	for q.Next() {
		err = q.Scan(&ID, &UserID, &Sum, &ProccessedAt)
		if err != nil {
			return nil, err
		}
		ret = append(ret, entity.Withdraw{
			ID:           ID,
			UserID:       UserID,
			Sum:          Sum,
			ProccessedAt: ProccessedAt,
		})
	}
	return ret, nil
}
