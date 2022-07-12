package storage

import "time"

//
// TODO: add JSON and SQL specification
//
type OrderQueue struct {
	StorageDB StorageDBparam

	OrderID        uint32
	TimeIn         time.Time
	TimeLastStatus time.Time
	Status         string
}

//
// TODO: realize interface StorageDBobjects
//

func (u *OrderQueue) Create() (err error)     { return nil }
func (u *OrderQueue) Read() (err error)       { return nil }
func (u *OrderQueue) Update() (err error)     { return nil }
func (u *OrderQueue) Delete() (err error)     { return nil }
func (u *OrderQueue) ExistOrNot() (err error) { return nil }

// OrderQueue's functions
