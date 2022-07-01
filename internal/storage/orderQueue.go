package storage

import "time"

//
// TODO: add JSON and SQL specification
//
type OrderQueue struct {
	StorageDB StorageDBparam

	OrderId        uint32
	TimeIn         time.Time
	TimeLastStatus time.Time
	Status         string
}

//
// TODO: realize interface StorageDBobjects
//

func (u *OrderQueue) Create() (err error)
func (u *OrderQueue) Read() (err error)
func (u *OrderQueue) Update() (err error)
func (u *OrderQueue) Delete() (err error)
func (u *OrderQueue) ExistOrNot() (err error)

// OrderQueue's functions
