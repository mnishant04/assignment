package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"unique;not_null"`
	Password string    `json:"password"`
	Address  string    `json:"address"`
	Product  []Product `json:"products,omitempty" gorm:"foreignKey:UserID"`
	LognActivity []LoginActivity `json:"login_activity,omitempty" gorm:"foreignKey:UserID"`
}

type LoginActivity struct {
	gorm.Model
	LoginTime  time.Time `json:"loginTime"`
	LogoutTime time.Time `json:"logoutTime"`
	UserID     *uint
}
