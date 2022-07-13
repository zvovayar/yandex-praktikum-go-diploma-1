package storage

import "gorm.io/gorm"

//
// TODO: add JSON and SQL specification
//
type User struct {
	gorm.Model
	Login      string `json:"login" gorm:"login;unique"`
	PasswdHash string `json:"password" gorm:"password"`
}

//
// TODO: realize interface StorageDBobjects
//
// User's functions
func (u *User) GetBalance() (err error) { return nil }
