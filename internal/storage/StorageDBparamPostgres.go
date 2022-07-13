package storage

import (
	"fmt"

	_ "github.com/lib/pq"
	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type StorageDBparamPostgres struct {
	uridb string
	db    *gorm.DB
}

var GORMinterface StorageDBparamPostgres

//
// TODO: realize interface StorageDBparam
//
func (sdbp *StorageDBparamPostgres) GetDB() (dbx *gorm.DB, err error) {
	if sdbp.db != nil {
		return sdbp.db, nil
	}

	sdbp.uridb = config.ConfigCLS.DataBaseURI

	sdbp.db, err = gorm.Open(postgres.Open(sdbp.uridb), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "gorm_", // table name prefix, table for `User` would be `t_users`
		},
	})
	if err != nil {
		config.LoggerCLS.Info(err.Error())
		return nil, fmt.Errorf("can,t open database postgres: " + err.Error())
	}

	err = sdbp.db.AutoMigrate(&User{})
	if err != nil {
		config.LoggerCLS.Info(err.Error())
		return nil, fmt.Errorf("can,t migrate User database postgres: " + err.Error())
	}

	return sdbp.db, nil
}
