package storage

import "time"

//
// TODO: add JSON and SQL specification
//
type OrderLog struct {
	StorageDB StorageDBparam

	OrderId uint32
	Time    time.Time
	Message string
	Status  string
}

//
// TODO: realize interface StorageDBobjects
//

func (u *OrderLog) Create() (err error)
func (u *OrderLog) Read() (err error)
func (u *OrderLog) Update() (err error)
func (u *OrderLog) Delete() (err error)
func (u *OrderLog) ExistOrNot() (err error)

// OrderLog's functions
