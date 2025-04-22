package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:'user'"`
}

type Device struct {
	gorm.Model
	UserID      uint
	User        User
	DeviceType  string
	ContactInfo string
}
