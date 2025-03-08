package handlers

import (
	"coffee_an1kv/server"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetImageHandler(s *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		entityType := c.Param("entity_type")
		entityID := c.Param("id")
		filename := c.Param("filename")

		filePath := filepath.Join(s.ImageDir, entityType, entityID, filename)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
			return
		}

		c.File(filePath)
	}
}
