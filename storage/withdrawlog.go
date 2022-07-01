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

func (u *WithdrawLog) Create() (err error)
func (u *WithdrawLog) Read() (err error)
func (u *WithdrawLog) Update() (err error)
func (u *WithdrawLog) Delete() (err error)
func (u *WithdrawLog) ExistOrNot() (err error)

// WithdrawLog's functions
