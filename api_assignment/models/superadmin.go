package models

import "gorm.io/gorm"

type SuperAdmin struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique;not_null"`
	Password string `json:"password"`
}
