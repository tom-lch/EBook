package gorm

import (
	"admin/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(new(models.SysConfig))
	if err != nil {
		return err
	}

}
