package db

import (
	"fmt"
	"log"

	"github.com/DomEscobar/erp-dev-bench/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto-migrate the schema
	err = DB.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	fmt.Println("Database connection established and migrated")
}

// InitTestDB initializes an in-memory SQLite database for testing
func InitTestDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect test database: %v", err)
	}

	// Auto-migrate the schema
	err = DB.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatalf("failed to migrate test database: %v", err)
	}

	fmt.Println("Test database initialized")
}

func Close() {
	sqlDB, err := DB.DB()
	if err == nil {
		sqlDB.Close()
	}
}
