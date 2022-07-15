package accrualclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
)

type Accrual struct {
	Address string
}

type Order struct {
	OrderNumber string `json:"order"`
	Goods       []Good `json:"goods"`
}

type Good struct {
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

func (a *Accrual) RegisterOrder(oid string) (err error) {

	//
	// call POST /api/orders
	// description: good name and price randomly
	//

	// make Order object
	order := Order{
		OrderNumber: oid,
		Goods: []Good{
			{Description: "Пеленальный столик Bork", Price: 123.45}, // always fix description and price
		},
	}

	body, err := json.Marshal(order)
	if err != nil {
		config.LoggerCLS.Sugar().Infow("%v", err)
		return err
	}

	config.LoggerCLS.Sugar().Debugf("order=%v", order)
	config.LoggerCLS.Sugar().Debugf("body=%v", string(body))

	var url = fmt.Sprintf("%v/api/orders",
		a.Address)

	config.LoggerCLS.Sugar().Debugf(url)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		config.LoggerCLS.Sugar().Error(err.Error())
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	config.LoggerCLS.Sugar().Debugf("request: %v\n", request)

	client := &http.Client{}

	// отправляем запрос
	resp, err := client.Do(request)
	if err != nil {
		config.LoggerCLS.Sugar().Error(err.Error())
		return err
	}
	defer resp.Body.Close()
	config.LoggerCLS.Sugar().Debug(resp.Status)

	if resp.StatusCode > 299 && resp.StatusCode != 409 {
		return fmt.Errorf("Accrual return: %v", resp)
	}
	return nil
}

func (a *Accrual) GetOrderStatus(onumber string) (status string, accrual float32, err error) {
	//
	// call GET /api/orders/{number}
	//

	config.LoggerCLS.Sugar().Debugf("get order status order=%v", onumber)

	var url = fmt.Sprintf("%v/api/orders/%v",
		a.Address, onumber)

	config.LoggerCLS.Sugar().Debugf(url)

	request, err := http.NewRequest("GET", url, bytes.NewBuffer(make([]byte, 0)))
	if err != nil {
		config.LoggerCLS.Sugar().Error(err.Error())
		return "", 0, err
	}
	request.Header.Set("Content-Type", "application/json")

	config.LoggerCLS.Sugar().Debugf("request: %v\n", request)

	client := &http.Client{}

	// отправляем запрос
	resp, err := client.Do(request)
	if err != nil {
		config.LoggerCLS.Sugar().Error(err.Error())
		return "", 0, err
	}
	defer resp.Body.Close()
	config.LoggerCLS.Sugar().Debug(resp.Status)

	if resp.StatusCode > 299 {
		return "", 0, fmt.Errorf("Accrual return: %v", resp)
	}

	type OrderFromAccrual struct {
		Order   string  `json:"order"`
		Status  string  `json:"status"`
		Accrual float32 `json:"accrual"`
	}

	var orderFromAccrual OrderFromAccrual
	// b := make([]byte, 10000)

	// decode orderstring
	// n, _ := resp.Body.Read(b)
	// config.LoggerCLS.Debug(fmt.Sprintf("body size:%d returned from accrual:%v", n, string(b)))

	if err := json.NewDecoder(resp.Body).Decode(&orderFromAccrual); err != nil {
		config.LoggerCLS.Sugar().Error(err.Error())
		return "", 0, err
	}

	config.LoggerCLS.Debug(fmt.Sprintf("order returned from accrual:%v", orderFromAccrual))
	return orderFromAccrual.Status, orderFromAccrual.Accrual, nil
}
