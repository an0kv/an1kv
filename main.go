package main

import (
	"coffee_an1kv/handlers"
	"coffee_an1kv/server"

	"github.com/gin-gonic/gin"
)

func main() {
	//gin
	r := gin.Default()
	s := server.NewServer()

	//роуты
	r.GET("/getusers", handlers.GetUsersHandler(s))
	r.POST("/addusers", handlers.AddUserHandler(s))
	r.GET("/categories", handlers.GetCategoriesHandler)
	r.GET("/products", handlers.GetProductsHandler)
	r.GET("/categories/:id", handlers.GetCategoryByID)
	r.GET("/categories/:id/products", handlers.GetProductsByCategoryID)

	//запуск сервера
	r.Run()
}
