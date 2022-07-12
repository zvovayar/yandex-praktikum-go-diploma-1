package storage

//
// TODO: add JSON and SQL specification
//
type Order struct {
	StorageDB StorageDBparam

	Id      uint32
	Accrual uint
	UserId  uint32
	Status  string
}

//
// TODO: realize interface StorageDBobjects
//

func (u *Order) Create() (err error)     { return nil }
func (u *Order) Read() (err error)       { return nil }
func (u *Order) Update() (err error)     { return nil }
func (u *Order) Delete() (err error)     { return nil }
func (u *Order) ExistOrNot() (err error) { return nil }

// Order's functions
