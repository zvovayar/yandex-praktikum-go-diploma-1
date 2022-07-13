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

//
// TODO: add JSON and SQL descriptions
//
type Order struct {
	OrderNumber string `json:"order"`
	Goods       []Good `json:"goods"`
}

//
// TODO: add JSON and SQL descriptions
//
type Good struct {
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

func (a *Accrual) RegisterOrder(oid string) (err error) {

	//
	// TODO: call POST /api/orders
	// description: good name and price randomly
	//

	// make Order object
	order := Order{
		OrderNumber: oid,
		Goods: []Good{
			{Description: "Пеленальный столик Bork", Price: 123.45},
		},
	}

	body, err := json.Marshal(order)
	if err != nil {
		config.LoggerCLS.Sugar().Infow("%v", err)
		return err
	}

	config.LoggerCLS.Sugar().Debugf("order=%v", order)
	config.LoggerCLS.Sugar().Debugf("body=%v", string(body))

	var url = fmt.Sprintf("http://%v/api/orders",
		a.Address)

	config.LoggerCLS.Sugar().Debugf(url)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		// обработаем ошибку
		config.LoggerCLS.Sugar().Error(err.Error())
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	config.LoggerCLS.Sugar().Debugf("request: %v\n", request)

	client := &http.Client{}

	// отправляем запрос
	resp, err := client.Do(request)
	if err != nil {
		// обработаем ошибку
		config.LoggerCLS.Sugar().Error(err.Error())
		return err
	}
	defer resp.Body.Close()
	config.LoggerCLS.Sugar().Debug(resp.Status)

	if resp.StatusCode > 299 {
		return fmt.Errorf("Accrual return: %v", resp)
	}
	return nil
}

func (a *Accrual) GetOrderStatus(onumber string) (status string, accrual float32, err error) {
	//
	// TODO: call GET /api/orders/{number}
	//
	return "", 0, nil
}
