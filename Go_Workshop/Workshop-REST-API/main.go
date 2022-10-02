package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rest-api/model"
)

func main() {
	r := gin.New()

	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.Books)
	})

	r.POST("/books", func(c *gin.Context) {
		var book model.Book

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		model.Books = append(model.Books, book)
		c.JSON(http.StatusCreated, book)
	})

	r.DELETE("/books/:id", func(c *gin.Context) {
		id := c.Param("id")

		for i, a := range model.Books {
			if a.ID == id {
				model.Books = append(model.Books[:i], model.Books[i+1:]...)
				break
			}
		}

		c.Status(http.StatusNoContent)
	})

	r.Run()
}
