package main

import (
	"coffee_an1kv/handlers"
	"coffee_an1kv/server"

	"github.com/gin-gonic/gin"
)

func main() {
	//gin
	r := gin.Default()
	s, err := server.NewServer("./images")
	if err != nil {
		panic("Server initialization error: " + err.Error())
	}

	//роуты
	r.GET("/getusers", handlers.GetUsersHandler(s))
	r.POST("/addusers", handlers.AddUserHandler(s))
	r.GET("/categories", handlers.GetCategoriesHandler)
	r.GET("/products", handlers.GetProductsHandler)
	r.GET("/categories/:id", handlers.GetCategoryByID)
	r.GET("/categories/:id/products", handlers.GetProductsByCategoryID)
	r.GET("/images/:filename", handlers.GetImageHandler(s))

	//запуск сервера
	r.Run()
}
