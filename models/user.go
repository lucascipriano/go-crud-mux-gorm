package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name     string `gorm:"size:64"`
	Email    string `gorm:"size:100"`
	Password string `gorm:"size:255"`
}
