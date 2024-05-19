package controller

import (
	"fmt"
	"go-app/initializers"
	"go-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var body struct {
		Name        string `json:"name"`
		Author      string `json:"author"`
		PublishYear int    `json:"publish_year"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		UserID:      1,
		Name:        body.Name,
		Author:      body.Author,
		PublishYear: body.PublishYear,
	}

	if err := initializers.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	message := fmt.Sprintf("Create book '%s' successfully", body.Name)
	c.JSON(http.StatusOK, gin.H{"message": message})
}
