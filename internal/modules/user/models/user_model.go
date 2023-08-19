package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"varchar:191"`
	Email    string `gorm:"varchar:191"`
	Password string `gorm:"varchar:191"`
}
