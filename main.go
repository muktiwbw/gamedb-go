package main

import "github.com/gin-gonic/gin"

func main() {
	// Initialize the routing API
	r := gin.Default()

	// Setting up routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Run the server
	r.Run() // listen and serve on localhost:8080
}
