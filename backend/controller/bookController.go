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

func ListBooks(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var books []models.Book
	result := initializers.DB.Where("user_id = ?", userID).Find(&books)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving books"})
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No books found for this user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func UpdateBook(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var bookUpdate struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Author      string `json:"author"`
		PublishYear int    `json:"publish_year"`
	}

	if err := c.ShouldBindJSON(&bookUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	var book models.Book
	result := initializers.DB.Where("user_id = ? AND id = ?", userID, bookUpdate.ID).First(&book)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving book"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not Found"})
		return
	}

	book.Name = bookUpdate.Name
	book.Author = bookUpdate.Author
	book.PublishYear = bookUpdate.PublishYear

	result = initializers.DB.Save(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": book})

}

func DeleteBook(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	id := c.Param("id")

	var book models.Book
	if err := initializers.DB.First(&book, book.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if book.UserID != uint(userID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	result := initializers.DB.Delete(&models.Book{}, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete Successfully"})

}
