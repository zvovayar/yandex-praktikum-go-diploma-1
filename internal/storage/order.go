package storage

import "gorm.io/gorm"

//
// TODO: add JSON and SQL specification
//
type Order struct {
	gorm.Model
	OrderNumber string `gorm:"unique"`
	Accrual     uint
	UserID      uint
	Status      string
}
