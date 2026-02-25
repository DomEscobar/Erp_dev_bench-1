package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DomEscobar/erp-dev-bench/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handlers.HealthCheck(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
}

func TestGetItems(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handlers.GetItems(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateItem(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/api/v1/items", strings.NewReader(`{"name":"Test","price":10.00,"sku":"TEST001"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	handlers.CreateItem(c)

	assert.Equal(t, http.StatusCreated, w.Code)
}
