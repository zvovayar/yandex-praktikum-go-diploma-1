package storage

import (
	"fmt"

	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
	"gorm.io/gorm"
)

type Withdraw struct {
	gorm.Model
	OrderNumber     string  `json:"order" gorm:"unique"`
	AccrualWithdraw float32 `json:"sum"`
	UserID          uint
	Status          string
}

func (w *Withdraw) CheckNewAndSave(uid uint) (st string, err error) {

	db, err := GORMinterface.GetDB()
	if err != nil {
		return "DBerror", err
	}

	// check balance
	// sum accrual in orders - sum withdrawals

	var count int64
	// count orders
	tx := db.Model(&Order{}).Where("user_id = ?", uid).Count(&count)
	if tx.Error != nil {
		return "DBerror", tx.Error
	}
	var sumOrders float32
	if count == 0 {
		sumOrders = 0
	} else {
		tx = db.Raw("SELECT SUM(accrual) FROM gorm_orders WHERE user_id = ?",
			uid).Scan(&sumOrders)
		if tx.RowsAffected == 0 {
			sumOrders = 0
		} else if tx.Error != nil {
			return "DBerror", tx.Error
		}
	}

	// count withdraws
	tx = db.Model(&Withdraw{}).Where("user_id = ?", uid).Count(&count)
	if tx.Error != nil {
		return "DBerror", tx.Error
	}
	var sumWithdraws float32
	if count == 0 {
		sumWithdraws = 0
	} else {
		tx = db.Raw("SELECT SUM(accrual_withdraw) FROM gorm_withdraws WHERE user_id = ?",
			uid).Scan(&sumWithdraws)
		if tx.RowsAffected == 0 {
			sumWithdraws = 0
		} else if tx.Error != nil {
			return "DBerror", tx.Error
		}
	}

	// if balance to small return error
	if w.AccrualWithdraw > sumOrders-sumWithdraws {
		return "Few", fmt.Errorf("w.AccrualWithdraw=%v > sumOrders=%v - sumWithdraws=%v",
			w.AccrualWithdraw, sumOrders, sumWithdraws)
	}

	// all is OK register withdraw
	config.LoggerCLS.Debug(fmt.Sprintf("w.AccrualWithdraw=%v, sumOrders=%v, sumWithdraws=%v",
		w.AccrualWithdraw, sumOrders, sumWithdraws))

	tx = db.Create(&w)
	if tx.Error != nil {
		return "DBerror", tx.Error
	}

	config.LoggerCLS.Sugar().Debugf("new withdraw registered successfuly: %v", w)

	return "OKRegistered", nil
}

func (w *Withdraw) GetByUser(uid uint) (wls []Withdraw, st string, err error) {

	db, err := GORMinterface.GetDB()
	if err != nil {
		return nil, "DBerror", err
	}

	tx := db.Order("created_at").Find(&wls, "user_id = ?", uid)
	if tx.Error != nil {
		return nil, "DBerror", tx.Error
	}

	return wls, "OK", nil
}
