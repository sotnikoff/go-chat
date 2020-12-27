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
	Port    int           `required:"true" default:"8080"`
	Timeout time.Duration `required:"true" default:"5s"`
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
	sig := <-quit
	log.Println("Shutting down server...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func handleRoutes(r *gin.Engine) {
	r.GET("/", indexPage)
	messages := r.Group("/messages")
	{
		messages.GET("/", getMessages)
		messages.POST("/", sendMessage)
		messages.PUT("/:id", editMessage)
		messages.DELETE("/:id", deleteMessage)
	}

	rooms := r.Group("/rooms")
	{
		rooms.GET("/", getRooms)
		rooms.POST("/", createRoom)
		rooms.POST("/:id/invite_user", inviteUser)
		rooms.POST("/:id/remove_user", removeUser)
		rooms.PUT("/:id", editRoom)
		rooms.DELETE("/:id", deleteRoom)
	}
}

func indexPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}

func getMessages(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Messages",
	})
}

func sendMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Send message",
	})
}

func editMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Send message",
	})
}

func deleteMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Message",
	})
}

func getRooms(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Room list",
	})
}

func editRoom(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Room edit",
	})
}

func createRoom(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Room create",
	})
}

func deleteRoom(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Room delete",
	})
}

func inviteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User invite to room",
	})
}

func removeUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User remove from room",
	})
}
