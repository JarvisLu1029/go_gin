package main

import (
	"go-app/controller"
	"go-app/initializers"

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
	}

	api := r.Group("/api")
	{
		api.POST("/posts", controller.PostCreate)
		api.POST("/book", controller.CreateBook)
	}

	r.Run()

}
