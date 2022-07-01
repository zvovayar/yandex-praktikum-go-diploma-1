package storage

type StorageDBobjects interface {
	Create() (err error)
	Read() (err error)
	Update() (err error)
	Delete() (err error)
	ExistOrNot() (err error)
}
