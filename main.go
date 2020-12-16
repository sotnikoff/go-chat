package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handleRoutes(r)
	r.Run()
}

func handleRoutes(r *gin.Engine) {
	r.GET("/", indexPage)
}

func indexPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}
