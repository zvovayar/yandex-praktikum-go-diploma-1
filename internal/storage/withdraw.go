package storage

//
// TODO: add JSON and SQL specification
//
type Withdraw struct {
	OrderNumber     string  `json:"order" gorm:"unique"`
	AccrualWithdraw float32 `json:"sum"`
	UserID          uint
	Status          string
}
