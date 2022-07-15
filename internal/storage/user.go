package storage

import (
	"strings"

	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login      string `json:"login" gorm:"login;unique"`
	PasswdHash string `json:"password" gorm:"password"`
}

// TODO: User's functions
func (u *User) GetBalance() (err error) { return nil }

func (u *User) CheckNewAndSave() (status string, err error) {

	db, err := GORMinterface.GetDB()
	if err != nil {
		return "DBerror", err
	}

	tx := db.Create(&u)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "duplicate key value violates unique constraint") {
			return "LoginBusy", tx.Error
		}
		return "DBerror", tx.Error
	}

	config.LoggerCLS.Sugar().Debugf("new user registered successfuly: %v", u)

	return "OKregistered", nil

}
