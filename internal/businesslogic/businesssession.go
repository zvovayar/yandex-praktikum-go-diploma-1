package businesslogic

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/osamingo/checkdigit"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/accrualclient"
	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/storage"
)

type BusinessSession struct {
	AccrualClient accrualclient.Accrual
}

type OrderForJSON struct {
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

func (bs *BusinessSession) LoadOrder(oc string, ulogin string) (status int, err error) {

	config.LoggerCLS.Debug(fmt.Sprintf("user %v load order number %v", ulogin, oc))

	// check Luhn algoritm
	if !checkdigit.NewLuhn().Verify(oc) {
		return 422, errors.New("order number is not valid by Luhn alogoritm: " + oc)
	}

	// check user exist?
	db, err := storage.GORMinterface.GetDB()

	if err != nil {
		return 500, err
	}

	var user storage.User
	tx := db.First(&user, "login = ?", ulogin)
	if tx.Error != nil {
		return 500, tx.Error
	}

	// register order in accrual
	err = bs.AccrualClient.RegisterOrder(oc)
	if err != nil {
		return 500, err
	}

	// save order in database
	var order storage.Order
	order.OrderNumber = oc

	var st string
	st, err = order.CheckNewAndSave(user.ID)

	switch st {
	case "DBerror":
		return 500, err
	case "OKloaded":
		return 200, nil
	case "LoadOtherUser":
		return 409, err
	case "OKnew":
		return 202, nil
	default:
		return 500, errors.New("unknown status returned by order saver")
	}
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

	var ordersForJSON []OrderForJSON
	ordersForJSON = make([]OrderForJSON, 0)

	for i := 0; i < len(orders); i++ {

		ordersForJSON = append(ordersForJSON, (OrderForJSON{
			Number:     orders[i].OrderNumber,
			Status:     orders[i].Status,
			Accrual:    orders[i].Accrual,
			UploadedAt: orders[i].CreatedAt,
		}))
	}

	config.LoggerCLS.Sugar().Debugf("orders in CLS dtabase with data from accrual for user:%v are:%v",
		ulogin, ordersForJSON)

	// make JSON
	jsonb, err = json.Marshal(ordersForJSON)
	if err != nil {
		return []byte(""), err
	}
	config.LoggerCLS.Sugar().Debugf("json orders in CLS dtabase with data from accrual for user:%v are:%v",
		ulogin, string(jsonb))
	return jsonb, nil
}

func (bs *BusinessSession) GetBalance(ulogin string) (jsonb []byte, err error) {

	config.LoggerCLS.Debug("get balance and make json for user: " + ulogin)

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

	// check balance
	// sum accrual in orders - sum withdrawals
	var count int64
	db.Model(&storage.Order{}).Where("user_id = ?", user.ID).Count(&count)
	var sumOrders float32
	if count == 0 {
		sumOrders = 0
	} else {
		tx = db.Raw("SELECT SUM(accrual) FROM gorm_orders WHERE user_id = ?",
			user.ID).Scan(&sumOrders)
		if tx.RowsAffected == 0 {
			sumOrders = 0
		} else if tx.Error != nil {
			return []byte(""), tx.Error
		}
	}

	db.Model(&storage.Withdraw{}).Where("user_id = ?", user.ID).Count(&count)
	var sumWithdraws float32
	if count == 0 {
		sumWithdraws = 0
	} else {
		tx = db.Raw("SELECT SUM(accrual_withdraw) FROM gorm_withdraws WHERE user_id = ?",
			user.ID).Scan(&sumWithdraws)
		if tx.RowsAffected == 0 {
			sumWithdraws = 0
		} else if tx.Error != nil {
			return []byte(""), tx.Error
		}
	}

	type Balance struct {
		Current   float32 `json:"current"`
		Withdrawn float32 `json:"withdrawn"`
	}

	b := Balance{
		Current:   sumOrders - sumWithdraws,
		Withdrawn: sumWithdraws,
	}

	config.LoggerCLS.Sugar().Debugf("balance in CLS dtabase for user:%v are:%v",
		ulogin, b)

	// make JSON
	jsonb, err = json.Marshal(b)
	if err != nil {
		return []byte(""), err
	}
	config.LoggerCLS.Sugar().Debugf("json balance in CLS dtabase for user:%v are:%v",
		ulogin, string(jsonb))

	return jsonb, nil
}

func (bs *BusinessSession) Withdraw(w storage.Withdraw, ulogin string) (status int, err error) {

	config.LoggerCLS.Debug(fmt.Sprintf("for user: %v withdraw register: %v ", ulogin, w))

	// check Luhn algoritm
	if !checkdigit.NewLuhn().Verify(w.OrderNumber) {
		return 422, errors.New("order number is not valid by Luhn alogoritm: " + w.OrderNumber)
	}
	// check user exist?
	db, err := storage.GORMinterface.GetDB()
	if err != nil {
		return 500, err
	}
	var user storage.User
	tx := db.First(&user, "login = ?", ulogin)
	if tx.Error != nil {
		return 500, tx.Error
	}
	w.UserID = user.ID

	// check balance
	// sum accrual in orders - sum withdrawals
	// save order in database

	var st string
	st, err = w.CheckNewAndSave(user.ID)

	switch st {
	case "DBerror":
		return 500, err
	case "OKRegistered":
		return 200, nil
	case "Few":
		return 402, err
	default:
		return 500, errors.New("unknown status returned by withdraw saver")
	}

}

func (bs *BusinessSession) GetWithdrawals(ulogin string) (jsonb []byte, err error) {

	config.LoggerCLS.Debug("get withdrawals and make json for user: " + ulogin)

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

	// select all withdrawals for user
	var withdrawals []storage.Withdraw
	tx = db.Order("created_at").Find(&withdrawals, "user_id = ?", user.ID)
	if tx.Error != nil {
		return []byte(""), tx.Error
	}

	config.LoggerCLS.Sugar().Debugf("withdrawals in CLS dtabase fo user:%v are:%v", ulogin, withdrawals)

	type WithdrawForJSON struct {
		Order       string    `json:"order"`
		Sum         float32   `json:"sum"`
		ProcessedAt time.Time `json:"processed_at"`
	}

	var withdrawalsForJSON []WithdrawForJSON
	withdrawalsForJSON = make([]WithdrawForJSON, 0)

	for i := 0; i < len(withdrawals); i++ {

		withdrawalsForJSON = append(withdrawalsForJSON, (WithdrawForJSON{
			Order:       withdrawals[i].OrderNumber,
			Sum:         withdrawals[i].AccrualWithdraw,
			ProcessedAt: withdrawals[i].CreatedAt,
		}))
	}

	config.LoggerCLS.Sugar().Debugf("withdrawals in CLS DB for user:%v are:%v",
		ulogin, withdrawalsForJSON)

	// make JSON
	jsonb, err = json.Marshal(withdrawalsForJSON)
	if err != nil {
		return []byte(""), err
	}
	config.LoggerCLS.Sugar().Debugf("json withdrawals in CLS DB for user:%v are:%v",
		ulogin, string(jsonb))
	return jsonb, nil
}

func (bs *BusinessSession) InfinityUpdateAllOrdersFromAccrual(dur time.Duration) (err error) {

	config.LoggerCLS.Debug("update all ordres statuses")
	// select all orders with not final statuses
	db, err := storage.GORMinterface.GetDB()
	if err != nil {
		return err
	}

	orders := make([]storage.Order, 0)

	go func() {
		for {
			tx := db.Where("status not in ?", []string{"INVALID", "PROCESSED"}).Find(&orders)
			if tx.Error != nil {
				return
			}
			if len(orders) > 0 {
				config.LoggerCLS.Debug(fmt.Sprintf("update all ordres statuses orders:%v", orders))
			}
			// check statuses and sums from accrual and update CLS DB
			var status string
			var accrual float32

			for i := 0; i < len(orders); i++ {

				status, accrual, err = bs.AccrualClient.GetOrderStatus(orders[i].OrderNumber)
				if err != nil {
					return
				}

				orders[i].Accrual = accrual
				orders[i].Status = status
				tx = db.Save(&orders[i])
				if tx.Error != nil {
					return
				}

			}
			<-time.After(dur)
		}
	}()

	return nil
}
