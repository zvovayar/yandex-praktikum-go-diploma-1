package storage

import (
	"time"
)

type OrderLog struct {
	OrderNumber string
	Time        time.Time
	Message     string
	Status      string
}
