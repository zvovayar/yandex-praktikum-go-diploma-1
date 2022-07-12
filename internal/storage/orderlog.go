package storage

import "time"

//
// TODO: add JSON and SQL specification
//
type OrderLog struct {
	StorageDB StorageDBparam

	OrderID uint32
	Time    time.Time
	Message string
	Status  string
}

//
// TODO: realize interface StorageDBobjects
//

func (u *OrderLog) Create() (err error)     { return nil }
func (u *OrderLog) Read() (err error)       { return nil }
func (u *OrderLog) Update() (err error)     { return nil }
func (u *OrderLog) Delete() (err error)     { return nil }
func (u *OrderLog) ExistOrNot() (err error) { return nil }

// OrderLog's functions
