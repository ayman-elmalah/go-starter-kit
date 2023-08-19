package database

import "gorm.io/gorm"

func Connection() *gorm.DB {
	return DB
}
