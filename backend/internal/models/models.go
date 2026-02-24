package models

import (
	"gorm.io/gorm"
)

// Base contains common fields for all models
type Base struct {
	gorm.Model
}
