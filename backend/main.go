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

	api := r.Group("/api")
	{
		api.POST("/posts", controller.PostCreate)
	}

	r.Run()

}
