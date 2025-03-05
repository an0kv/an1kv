package handlers

import (
	"coffee_an1kv/server"

	"github.com/gin-gonic/gin"
)

func GetUsersHandler(s *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, s.GetUsers())
	}
}
