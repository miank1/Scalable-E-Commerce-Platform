package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "list of products",
		})
	})

	r.Run(":8081")
}
