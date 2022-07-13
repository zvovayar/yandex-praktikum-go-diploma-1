package storage

import "gorm.io/gorm"

type StorageDBparam interface {
	GetDB() (db *gorm.DB, err error)
}
