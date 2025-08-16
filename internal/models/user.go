package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PhoneNumber string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
}
