package service

import "errors"

var (
	ErrOrderIsIncorrect      = errors.New("order number is incorrect")
	ErrOrderAlreadySave      = errors.New("order saved")
	ErrOrderSavedByOtherUser = errors.New("order saved by other user")
)
