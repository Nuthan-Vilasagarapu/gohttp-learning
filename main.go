package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})

	router.POST("/form", func(c *gin.Context) {
		name := c.Request.FormValue("s_name")
		c.JSON(200, gin.H{
			"message": "Thank you " + name + ", for filling this form",
		})
	})

	router.Run(":8000")
}
