package handlers

import (
	"coffee_an1kv/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetImageHandler(s *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		filename := c.Param("filename")

		filePath, err := s.GetImagePath(filename)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
			return
		}

		c.File(filePath)
	}
}
