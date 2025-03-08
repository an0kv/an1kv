package handlers

import (
	"coffee_an1kv/server"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImageHandler(s *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		entityType := c.Param("entity_type")
		entityID := c.Param("id")
		id, _ := strconv.Atoi(entityID)

		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File not found"})
			return
		}

		ext := filepath.Ext(file.Filename)
		newName := fmt.Sprintf("%s_%d%s", entityType, time.Now().UnixNano(), ext)
		dst := filepath.Join(s.ImageDir, entityType, entityID, newName)

		os.MkdirAll(filepath.Dir(dst), 0755)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Saving error"})
			return
		}

		switch entityType {
		case "categories":
			for i := range server.Categories {
				if server.Categories[i].ID == id {
					server.Categories[i].Image = newName
					break
				}
			}
		case "products":
			for i := range server.Products {
				if server.Products[i].ID == id {
					server.Products[i].Image = newName
					break
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Image added",
			"image":   newName,
		})
	}
}
