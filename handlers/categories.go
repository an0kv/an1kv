package handlers

import (
	"coffee_an1kv/server"

	"github.com/gin-gonic/gin"
)

func GetCategoriesHandler(c *gin.Context) {
	c.JSON(200, server.Categories)
}
