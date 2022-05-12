package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indranilbiswas1984/bookstore/controllers"
	"github.com/indranilbiswas1984/bookstore/models"
)

func main() {
	println("Welcome")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook) // new

	r.Run()

}
