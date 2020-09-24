package database

import (
	"admin/global"
	"admin/tools"
	"admin/tools/config"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type MySQL struct {
}


func (sql *MySQL)Setup() {
	var err error
	var db Database
	db = new(MySQL)
	global.Source = db.GetConnect()
	global.Driver = db.GetDriver()
	global.GDB, err = db.Open(db.GetDriver()+db.GetConnect())
	if err != nil {
		global.Logger.Fatal(tools.Red(db.GetDriver() + "connect error :"), err)
	} else {
		global.Logger.Info(tools.Green(db.GetDriver() + "connect success!"))
	}
	if global.GDB.Error != nil {
		global.Logger.Fatal(tools.Red(" database error :"),  global.GDB.Error)
	}
}

func (sql *MySQL)Open(dsn string) (db *gorm.DB, err error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
func (sql *MySQL)GetConnect() string {
	return config.DatabaseConfig.Source
}
func (sql *MySQL)GetDriver() string {
	return config.DatabaseConfig.Driver
}
