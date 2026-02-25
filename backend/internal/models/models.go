package models

import (
	"time"

	"gorm.io/gorm"
)

// Base contains common fields for all models with proper JSON tags
type Base struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
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
