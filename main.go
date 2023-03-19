package main

import (
	"net/http"

	"github.com/diegoall/gin-bookstore/controllers"
	"github.com/diegoall/gin-bookstore/models"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// RESTful API usually give the response in JSON format, but there are several types of response to the client.
		c.JSON(http.StatusOK, gin.H{"data": "Que mas pues"})
	})

	// Connect to database
	models.ConnectDatabase() // new

	// Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run(":9090")
}
