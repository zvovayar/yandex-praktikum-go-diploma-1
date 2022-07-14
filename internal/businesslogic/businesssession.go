package businesslogic

import (
	"errors"
	"fmt"

	"github.com/osamingo/checkdigit"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/accrualclient"
	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
	httpcs "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/httpserver/sessions"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/storage"
)

type BusinessSession struct {
	HTTPsession httpcs.CurrentSession
}

func (bs *BusinessSession) RegisterNewUser(u storage.User) (err error) {

	config.LoggerCLS.Debug("register new user " + u.Login)

	db, err := storage.GORMinterface.GetDB()

	if err != nil {
		return err
	}

	tx := db.Create(&u)
	if tx.Error != nil {
		return tx.Error
	}

	config.LoggerCLS.Sugar().Debugf("new user registered successfuly: %v", u)

	return nil
}
func (bs *BusinessSession) UserLogin(u storage.User) (err error) {

	config.LoggerCLS.Debug("login user " + u.Login)

	db, err := storage.GORMinterface.GetDB()

	if err != nil {
		return err
	}

	var user storage.User
	tx := db.First(&user, "login = ?", u.Login)
	if tx.Error != nil {
		return tx.Error
	}

	if user.PasswdHash != u.PasswdHash {
		return errors.New("password failed")
	}

	return nil
}

func (bs *BusinessSession) LoadOrder(oc string, ulogin string) (err error) {

	config.LoggerCLS.Debug(fmt.Sprintf("user %v load order number %v", ulogin, oc))

	// check Luhn algoritm
	if !checkdigit.NewLuhn().Verify(oc) {
		return errors.New("order number is not valid by Luhn alogoritm: " + oc)
	}

	// check user exist?
	db, err := storage.GORMinterface.GetDB()

	if err != nil {
		return err
	}

	var user storage.User
	tx := db.First(&user, "login = ?", ulogin)
	if tx.Error != nil {
		return tx.Error
	}

	// register order in accrual
	err = (&(accrualclient.Accrual{Address: config.ConfigCLS.AccrualSystemAddress})).RegisterOrder(oc)
	if err != nil {
		return err
	}

	// save order in database
	order := storage.Order{
		// Model:       gorm.Model{},
		OrderNumber: oc,
		// Accrual:     0,
		UserID: user.ID,
		// Status:      "",
	}

	tx = db.Create(&order)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (bs *BusinessSession) GetOrders() (json string, err error) {
	config.LoggerCLS.Debug("read orders and make json ")

	json = `[
		{
			"number": "9278923470",
			"status": "PROCESSED",
			"accrual": 500,
			"uploaded_at": "2020-12-10T15:15:45+03:00"
		},
		{
			"number": "12345678903",
			"status": "PROCESSING",
			"uploaded_at": "2020-12-10T15:12:01+03:00"
		},
		{
			"number": "346436439",
			"status": "INVALID",
			"uploaded_at": "2020-12-09T16:09:53+03:00"
		}
	]`

	return json, nil
}

func (bs *BusinessSession) GetBalance() (json string, err error) {
	config.LoggerCLS.Debug("read balance and make json ")

	json = `{
		"current": 500.5,
		"withdrawn": 42
	}`

	return json, nil
}

func (bs *BusinessSession) Withdraw(w storage.Withdraw) (err error) {
	config.LoggerCLS.Debug("withdraw register ")
	return nil
}

func (bs *BusinessSession) GetWithdrawals() (json string, err error) {
	config.LoggerCLS.Debug("read withdrawals balance and make json ")

	json = `[
		{
			"order": "2377225624",
			"sum": 500,
			"processed_at": "2020-12-09T16:09:57+03:00"
		}
	]`

	return json, nil
}

func (bs *BusinessSession) Buy() (err error) { return nil }
