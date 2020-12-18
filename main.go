package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"log"
)

// Specification ...
type Specification struct {
	Port int `required:"true"`
}

func main() {
	var s Specification
	err := envconfig.Process("chat", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := gin.Default()
	handleRoutes(r)

	err = r.Run(fmt.Sprint(":", s.Port))
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
