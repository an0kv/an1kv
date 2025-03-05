package handlers

import (
	"coffee_an1kv/server"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProductsByCategoryID(c *gin.Context) {
	id := c.Param("id")
	var categoryProducts []server.Product
	productsID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid category ID"})
		return
	}
	for _, product := range server.Products {
		if product.CategoryID == productsID {
			categoryProducts = append(categoryProducts, product)
		}
	}
	if len(categoryProducts) > 0 {
		c.JSON(200, categoryProducts)
	} else {
		c.JSON(404, gin.H{"error": "No products found for this category"})
	}
}
