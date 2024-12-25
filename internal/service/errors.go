package service

import "errors"

var (
	ErrOrderAlreadySave      = errors.New("order saved")
	ErrOrderSavedByOtherUser = errors.New("order saved by other user")
)
