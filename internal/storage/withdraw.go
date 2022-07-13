package storage

//
// TODO: add JSON and SQL specification
//
type Withdraw struct {
	OrderNumber     string `json:"order" gorm:"unique"`
	AccrualWithdraw uint   `json:"sum"`
	UserID          uint
	Status          string
}
