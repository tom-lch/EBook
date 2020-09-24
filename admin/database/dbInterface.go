package database

import "gorm.io/gorm"

type Database interface {
	Open(dsn string) (db *gorm.DB, err error)
	GetConnect() string
	GetDriver() string
	Setup()
}
