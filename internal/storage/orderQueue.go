package storage

import (
	"time"

	"gorm.io/gorm"
)

type OrderQueue struct {
	gorm.Model
	OrderNumber    string `gorm:"unique"`
	TimeLastStatus time.Time
	Status         string
}
