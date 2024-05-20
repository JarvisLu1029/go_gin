package main

import (
	"go-app/controller"
	"go-app/initializers"
	"go-app/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/register", controller.RegisterUser)
		auth.POST("/login", controller.LoginUser)
	}

	api := r.Group("/api")
	{
		api.Use(middleware.AuthMiddleware())
		api.POST("/posts", controller.PostCreate)

		api.POST("/books", controller.CreateBook)
		api.GET("/books", controller.ListBooks)
		// api.GET("/books/:id", controller.ShowBook)
		api.PUT("/books/:id", controller.UpdateBook)
		api.DELETE("/books/:id", controller.DeleteBook)
	}

	r.Run()

}
