package entities

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Name     string
}
