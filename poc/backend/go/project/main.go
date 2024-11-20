package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET(":word", func(c *gin.Context) {
		word := c.Param("word")
		c.JSON(200, gin.H{
			"message": word,
		})
	})
	server.Run(":8080")
}