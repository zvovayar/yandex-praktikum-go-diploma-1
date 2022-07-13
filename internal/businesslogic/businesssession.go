package businesslogic

import (
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

	return nil
}
func (bs *BusinessSession) UserLogin(u storage.User) (err error) {

	config.LoggerCLS.Debug("login user " + u.Login)
	return nil
}

func (bs *BusinessSession) LoadOrder(oc string) (err error) {
	config.LoggerCLS.Debug("load order " + oc)

	err = (&(accrualclient.Accrual{Address: config.ConfigCLS.AccrualSystemAddress})).RegisterOrder(oc)
	return err
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
