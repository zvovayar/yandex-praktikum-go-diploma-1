package storage

import "time"

//
// TODO: add JSON and SQL specification
//
type WithdrawLog struct {
	OrderNumber string
	Time        time.Time
	Message     string
	Status      string
}
