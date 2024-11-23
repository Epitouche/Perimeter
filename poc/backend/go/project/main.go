package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/:word", func(c *gin.Context) {
		word := c.Param("word")
		c.JSON(200, gin.H{
			"message": word,
		})
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
