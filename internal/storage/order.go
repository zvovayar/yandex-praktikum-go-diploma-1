package storage

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderNumber string `gorm:"unique"`
	Accrual     float32
	UserID      uint
	Status      string
}

func (o *Order) CheckNewAndSave(uid uint) (status string, err error) {

	db, err := GORMinterface.GetDB()
	if err != nil {
		return "DBerror", err
	}

	o.Status = "NEW"
	o.UserID = uid

	tx := db.Create(o)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "duplicate key value violates unique constraint") {

			var order Order
			tx := db.First(&order, "order_number = ?", o.OrderNumber)
			if tx.Error != nil {
				return "DBerror", tx.Error
			}
			if order.UserID == uid {
				return "OKloaded", nil
			}
			return "LoadOtherUser", errors.New("order number " + o.OrderNumber + " was loaded by other user")
		}
		return "DBerror", tx.Error
	}
	return "OKnew", nil
}

func (o *Order) GetByUser(uid int) (orders []Order, status string, err error) {

	// select order numbers for userid
	db, err := GORMinterface.GetDB()
	if err != nil {
		return nil, "DBerror", err
	}
	tx := db.Order("created_at").Find(&orders, "user_id = ?", uid)
	if tx.Error != nil {
		return nil, "DBerror", tx.Error
	}

	return orders, "OK", nil
}

func (o *Order) GetQueueForAccrualUpdate() (orders []Order, status string, err error) {

	// select order numbers for queue for accrual statuses update
	db, err := GORMinterface.GetDB()
	if err != nil {
		return nil, "DBerror", err
	}
	tx := db.Where("status not in ?", []string{"INVALID", "PROCESSED"}).Find(&orders)
	if tx.Error != nil {
		return nil, "DBerror", err
	}

	return orders, "OK", nil
}

func (o *Order) Save() (status string, err error) {

	db, err := GORMinterface.GetDB()
	if err != nil {
		return "DBerror", err
	}

	tx := db.Save(o)
	if tx.Error != nil {
		return "DBerror", tx.Error
	}

	return "OK", nil

}
