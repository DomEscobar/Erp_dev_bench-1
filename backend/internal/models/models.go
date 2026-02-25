package models

import (
	"gorm.io/gorm"
)

// Base contains common fields for all models
type Base struct {
	gorm.Model
}

// Item represents a product/item in the inventory
type Item struct {
	Base
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null"`
	SKU         string  `json:"sku" gorm:"unique"`
	CategoryID  uint    `json:"categoryId"`
}
