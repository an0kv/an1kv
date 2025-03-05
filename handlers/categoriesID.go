package handlers

import (
	"strconv"

	"coffee_an1kv/server"

	"github.com/gin-gonic/gin"
)

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	categoryID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid category ID"})
		return
	}

	for _, category := range server.Categories {
		if category.ID == categoryID {
			c.JSON(200, category)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Category not found"})
}
