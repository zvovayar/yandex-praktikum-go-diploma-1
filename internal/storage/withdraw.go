package storage

//
// TODO: add JSON and SQL specification
//
type Withdraw struct {
	StorageDB StorageDBparam

	OrderId         uint32
	AccrualWithdraw uint
	UserId          uint32
	Status          string
}

//
// TODO: realize interface StorageDBobjects
//

func (u *Withdraw) Create() (err error)
func (u *Withdraw) Read() (err error)
func (u *Withdraw) Update() (err error)
func (u *Withdraw) Delete() (err error)
func (u *Withdraw) ExistOrNot() (err error)

// Withdraw's functions