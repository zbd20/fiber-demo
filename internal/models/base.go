package models

import (
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	tables := []interface{}{&User{}, &Login{}}
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(tables...)
	if err != nil {
		return err
	}

	return nil
}
