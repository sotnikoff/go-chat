package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	handleRoutes(r)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func handleRoutes(r *gin.Engine) {
	r.GET("/", indexPage)
}

func indexPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}
