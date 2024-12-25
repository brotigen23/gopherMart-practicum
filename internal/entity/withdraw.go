package entity

import "time"

type Withdraw struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Sum          float32   `json:"sum"`
	ProccessedAt time.Time `json:"processed_at"`
}
