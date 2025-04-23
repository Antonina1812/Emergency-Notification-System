package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string
	Role     string
}

type Device struct {
	gorm.Model
	UserID      uint
	User        User
	DeviceType  string
	ContactInfo string
}

type Message struct {
	gorm.Model
	Text string
	// Template bool
	CreatedBy uint
	UserID    uint
	User      User `gorm:"foreignKey:UserID"`
}

type MessageHistory struct {
	gorm.Model
	MessageID uint
	Message   Message
	DeviceID  uint
	Device    Device
	Status    string
}
