package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DomEscobar/erp-dev-bench/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestItems_GetItems(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handlers.GetItems(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestItems_CreateItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	uniqueSKU := fmt.Sprintf("TEST-%d", time.Now().UnixNano())
	body := fmt.Sprintf(`{"name":"Test","price":10.00,"sku":"%s"}`, uniqueSKU)
	c.Request = httptest.NewRequest("POST", "/api/v1/items", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")
	handlers.CreateItem(c)
	assert.Equal(t, http.StatusCreated, w.Code)
}
