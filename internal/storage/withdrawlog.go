package storage

import "time"

type WithdrawLog struct {
	OrderNumber string
	Time        time.Time
	Message     string
	Status      string
}
