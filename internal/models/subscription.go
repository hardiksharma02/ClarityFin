package models

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	Name   string
	Amount float64
	UserID uint // Foreign key to the User model
}
