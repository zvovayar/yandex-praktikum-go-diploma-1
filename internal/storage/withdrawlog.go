package storage

import "time"

//
// TODO: add JSON and SQL specification
//
type WithdrawLog struct {
	StorageDB StorageDBparam

	OrderId uint32
	Time    time.Time
	Message string
	Status  string
}

//
// TODO: realize interface StorageDBobjects
//

func (u *WithdrawLog) Create() (err error)     { return nil }
func (u *WithdrawLog) Read() (err error)       { return nil }
func (u *WithdrawLog) Update() (err error)     { return nil }
func (u *WithdrawLog) Delete() (err error)     { return nil }
func (u *WithdrawLog) ExistOrNot() (err error) { return nil }

// WithdrawLog's functions
