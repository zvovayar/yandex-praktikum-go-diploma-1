package storage

import (
	"time"
)

//
// TODO: add JSON and SQL specification
//
type OrderLog struct {
	OrderNumber string
	Time        time.Time
	Message     string
	Status      string
}
