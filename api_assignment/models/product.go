package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Image  string  `json:"image"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	UserID uint
}
