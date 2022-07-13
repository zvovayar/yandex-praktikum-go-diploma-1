package storage

import (
	"time"

	"gorm.io/gorm"
)

//
// TODO: add JSON and SQL specification
//
type WithdrawQueue struct {
	gorm.Model
	OrderNumber    string `gorm:"unique"`
	TimeLastStatus time.Time
	Status         string
}
