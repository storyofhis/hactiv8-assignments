package entity

import "gorm.io/gorm"

var DB *gorm.DB

type User struct {
	gorm.Model
	Email    string
	Username string
	Name     string
	Age      string
}
