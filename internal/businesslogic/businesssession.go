package businesslogic

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/osamingo/checkdigit"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/accrualclient"
	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
	httpcs "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/httpserver/sessions"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/storage"
)

type BusinessSession struct {
	HTTPsession httpcs.CurrentSession
}

type OrderForJson struct {
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    float32   `json:"accrual"`
	UploadedAt time.Time `json:"uploaded_at"`
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
		Status: "NEW",
	}

	tx = db.Create(&order)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (bs *BusinessSession) GetOrders(ulogin string) (jsonb []byte, err error) {

	config.LoggerCLS.Debug("read orders and make json for user: " + ulogin)

	// check user exist?
	db, err := storage.GORMinterface.GetDB()
	if err != nil {
		return []byte(""), err
	}
	var user storage.User
	tx := db.First(&user, "login = ?", ulogin)
	if tx.Error != nil {
		return []byte(""), tx.Error
	}

	// select order numbers for userid
	var orders []storage.Order
	tx = db.Find(&orders, "user_id = ?", user.ID)
	if tx.Error != nil {
		return []byte(""), tx.Error
	}

	config.LoggerCLS.Sugar().Debugf("orders in CLS dtabase fo user:%v are:%v", ulogin, orders)

	// get accrual statuses for orders from CLS database
	var status string
	var accrual float32

	var ordersForJson []OrderForJson
	ordersForJson = make([]OrderForJson, 0)

	for i := 0; i < len(orders); i++ {

		status, accrual, err = (&(accrualclient.Accrual{
			Address: config.ConfigCLS.AccrualSystemAddress,
		})).GetOrderStatus(orders[i].OrderNumber)
		if err != nil {
			return []byte(""), err
		}

		ordersForJson = append(ordersForJson, (OrderForJson{
			Number:     orders[i].OrderNumber,
			Status:     status,
			Accrual:    accrual,
			UploadedAt: orders[i].CreatedAt,
		}))
	}

	config.LoggerCLS.Sugar().Debugf("orders in CLS dtabase with data from accrual for user:%v are:%v",
		ulogin, ordersForJson)

	// make JSON
	jsonb, err = json.Marshal(ordersForJson)
	if err != nil {
		return []byte(""), err
	}
	config.LoggerCLS.Sugar().Debugf("json orders in CLS dtabase with data from accrual for user:%v are:%v",
		ulogin, string(jsonb))
	return jsonb, nil
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
