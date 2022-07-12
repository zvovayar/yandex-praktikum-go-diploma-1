package storage

import "time"

//
// TODO: add JSON and SQL specification
//
type WithdrawQueue struct {
	StorageDB StorageDBparam

	OrderID        uint32
	TimeIn         time.Time
	TimeLastStatus time.Time
	Status         string
}

//
// TODO: realize interface StorageDBobjects
//

func (u *WithdrawQueue) Create() (err error)     { return nil }
func (u *WithdrawQueue) Read() (err error)       { return nil }
func (u *WithdrawQueue) Update() (err error)     { return nil }
func (u *WithdrawQueue) Delete() (err error)     { return nil }
func (u *WithdrawQueue) ExistOrNot() (err error) { return nil }

// WithdrawQueue's functions
