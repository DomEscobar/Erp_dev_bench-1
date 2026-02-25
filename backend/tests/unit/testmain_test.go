package handlers

import (
	"os"
	"testing"

	"github.com/DomEscobar/erp-dev-bench/internal/db"
)

func TestMain(m *testing.M) {
	// Initialize in-memory DB for unit tests
	db.InitTestDB()
	defer db.Close()

	os.Exit(m.Run())
}
