package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Specification ...
type Specification struct {
	Port    int `required:"true" default:"8080"`
	Timeout int `required:"true" default:"5"`
}

func main() {
	var s Specification
	err := envconfig.Process("chat", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := gin.Default()

	handleRoutes(r)

	srv := &http.Server{
		Addr:    fmt.Sprint(":", s.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.Timeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func handleRoutes(r *gin.Engine) {
	r.GET("/", indexPage)
}

func indexPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}
