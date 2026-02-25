package handlers

import (
	"net/http"

	"github.com/DomEscobar/erp-dev-bench/internal/db"
	"github.com/DomEscobar/erp-dev-bench/internal/models"
	"github.com/gin-gonic/gin"
)

// GetItems handles GET /api/v1/items - lists all items
func GetItems(c *gin.Context) {
	var items []models.Item
	result := db.DB.Find(&items)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetItem handles GET /api/v1/items/:id - gets a single item
func GetItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	result := db.DB.First(&item, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

// CreateItem handles POST /api/v1/items - creates a new item
func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Create(&item)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// UpdateItem handles PUT /api/v1/items/:id - updates an item
func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	// Check if item exists
	if err := db.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var update models.Item
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	item.Name = update.Name
	item.Description = update.Description
	item.Price = update.Price
	item.SKU = update.SKU
	item.CategoryID = update.CategoryID

	if err := db.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteItem handles DELETE /api/v1/items/:id - deletes an item
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	// Check if item exists
	if err := db.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := db.DB.Delete(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
