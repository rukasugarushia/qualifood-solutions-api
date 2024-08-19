package handlers

import (
	"net/http"
	"qualifood-solutions-api/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

var bookService = domain.NewBookService()

func GetAllBooks(c *gin.Context) {
	books, _ := bookService.GetAll()
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, _ := bookService.GetByID(id)
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookService.Create(book)
	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookService.Update(id, book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	bookService.Delete(id)
	c.JSON(http.StatusNoContent, nil)
}
