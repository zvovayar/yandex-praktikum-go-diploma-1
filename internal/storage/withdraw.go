package storage

//
// TODO: add JSON and SQL specification
//
type Withdraw struct {
	StorageDB StorageDBparam

	OrderID         uint32
	AccrualWithdraw uint
	UserID          uint32
	Status          string
}

//
// TODO: realize interface StorageDBobjects
//

func (u *Withdraw) Create() (err error)     { return nil }
func (u *Withdraw) Read() (err error)       { return nil }
func (u *Withdraw) Update() (err error)     { return nil }
func (u *Withdraw) Delete() (err error)     { return nil }
func (u *Withdraw) ExistOrNot() (err error) { return nil }

// Withdraw's functions
