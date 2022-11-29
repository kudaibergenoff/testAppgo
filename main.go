package main

import (
	"github.com/gin-gonic/gin"
	"testApp/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Index endpoint",
		})
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Test endpoint",
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
