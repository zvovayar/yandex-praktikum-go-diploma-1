package storage

import "time"

//
// TODO: add JSON and SQL specification
//
type WithdrawQueue struct {
	StorageDB StorageDBparam

	OrderId        uint32
	TimeIn         time.Time
	TimeLastStatus time.Time
	Status         string
}

//
// TODO: realize interface StorageDBobjects
//

func (u *WithdrawQueue) Create() (err error)
func (u *WithdrawQueue) Read() (err error)
func (u *WithdrawQueue) Update() (err error)
func (u *WithdrawQueue) Delete() (err error)
func (u *WithdrawQueue) ExistOrNot() (err error)

// WithdrawQueue's functions
