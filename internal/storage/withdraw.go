package storage

import "gorm.io/gorm"

type Withdraw struct {
	gorm.Model
	OrderNumber     string  `json:"order" gorm:"unique"`
	AccrualWithdraw float32 `json:"sum"`
	UserID          uint
	Status          string
}
