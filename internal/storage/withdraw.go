package storage

import "gorm.io/gorm"

//
// TODO: add JSON and SQL specification
//
type Withdraw struct {
	gorm.Model
	OrderNumber     string  `json:"order" gorm:"unique"`
	AccrualWithdraw float32 `json:"sum"`
	UserID          uint
	Status          string
}
