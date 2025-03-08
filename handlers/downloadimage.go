package handlers

import (
	"coffee_an1kv/server"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadImageHandler(s *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Файл не найден в запросе"})
			return
		}
		ext := filepath.Ext(file.Filename)
		allowedExtensions := map[string]bool{".jpg": true, ".png": true, ".jpeg": true}
		if !allowedExtensions[ext] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Недопустимый формат файла"})
			return
		}

		dst := filepath.Join(s.ImageDir, file.Filename)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить файл"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Файл успешно загружен",
			"path":    dst,
		})
	}
}
