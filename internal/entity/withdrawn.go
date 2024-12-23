package entity

import "time"

type Withdrawn struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Sum          float64   `json:"sum"`
	ProccessedAt time.Time `json:"processed_at"`
}
