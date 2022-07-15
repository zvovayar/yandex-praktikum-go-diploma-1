package storage

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login      string `json:"login" gorm:"login;unique"`
	PasswdHash string `json:"password" gorm:"password"`
}

// TODO: User's functions
func (u *User) GetBalance() (err error) { return nil }
