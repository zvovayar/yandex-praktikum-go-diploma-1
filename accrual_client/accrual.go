package accrual_client

type Accrual struct {
	Address string
}

//
// TODO: add JSON and SQL descriptions
//
type Order struct {
	OrderNumber string
	Goods       []Good
}

//
// TODO: add JSON and SQL descriptions
//
type Good struct {
	Description string
	Price       int
}

func (a *Accrual) RegisterOrder(o Order) (err error) {

	//
	// TODO: call POST /api/orders
	//
	return nil
}

func (a *Accrual) GetOrderStatus(onumber string) (status string, err error) {
	//
	// TODO: call GET /api/orders/{number}
	//
	return "", nil
}
