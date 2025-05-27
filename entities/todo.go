package entities

import "gorm.io/gorm"

type Todos struct {
	gorm.Model
	UsersID   uint
	Users     Users
	Todo      string `json:"Todo"`
	Category  string `json:"Category"`
	IsChecked bool   `json:"IsChecked"`
}
