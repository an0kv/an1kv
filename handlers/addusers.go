package handlers

import (
	"coffee_an1kv/server"

	"github.com/gin-gonic/gin"
)

func AddUserHandler(s *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user server.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
		c.JSON(201, s.AddUser(user))
	}
}
