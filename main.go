package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/coba", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Halo Dunia",
		})
	})

	r.Run(":8080")
}
