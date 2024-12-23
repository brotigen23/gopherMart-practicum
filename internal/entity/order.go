package entity

import (
	"time"
)

type Order struct {
	ID         int `json:"order"`
	UserID     int
	Order      string
	UploadedAt time.Time
}
