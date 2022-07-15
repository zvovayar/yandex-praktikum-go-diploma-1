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
func (u *User) GetBalance() (sumOrders float32, sumWithdraws float32, status string, err error) {

	db, err := GORMinterface.GetDB()
	if err != nil {
		return 0, 0, "DBerror", err
	}

	// sum accrual in orders - sum withdrawals

	var count int64
	tx := db.Model(&Order{}).Where("user_id = ?", u.ID).Count(&count)
	if tx.Error != nil {
		return 0, 0, "DBerror", tx.Error
	}

	if count == 0 {
		sumOrders = 0
	} else {
		tx = db.Raw("SELECT SUM(accrual) FROM gorm_orders WHERE user_id = ?",
			u.ID).Scan(&sumOrders)
		if tx.RowsAffected == 0 {
			sumOrders = 0
		} else if tx.Error != nil {
			return 0, 0, "DBerror", tx.Error
		}
	}

	tx = db.Model(&Withdraw{}).Where("user_id = ?", u.ID).Count(&count)
	if tx.Error != nil {
		return 0, 0, "DBerror", tx.Error
	}

	if count == 0 {
		sumWithdraws = 0
	} else {
		tx = db.Raw("SELECT SUM(accrual_withdraw) FROM gorm_withdraws WHERE user_id = ?",
			u.ID).Scan(&sumWithdraws)
		if tx.RowsAffected == 0 {
			sumWithdraws = 0
		} else if tx.Error != nil {
			return 0, 0, "DBerror", tx.Error
		}
	}

	return sumOrders, sumWithdraws, "OK", nil
}

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
