package storage

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderNumber string `gorm:"unique"`
	Accrual     float32
	UserID      uint
	Status      string
}
