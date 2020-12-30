package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/sotnikoff/go-chat/internal/messages"
	"github.com/sotnikoff/go-chat/internal/rooms"
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
	msgGroup := r.Group("/messages")
	{
		msgGroup.GET("/", messages.GetMessages)
		msgGroup.POST("/", messages.SendMessage)
		msgGroup.PUT("/:id", messages.EditMessage)
		msgGroup.DELETE("/:id", messages.DeleteMessage)
	}

	roomsGroup := r.Group("/rooms")
	{
		roomsGroup.GET("/", rooms.GetRooms)
		roomsGroup.POST("/", rooms.CreateRoom)
		roomsGroup.POST("/:id/invite_user", rooms.InviteUser)
		roomsGroup.POST("/:id/remove_user", rooms.RemoveUser)
		roomsGroup.PUT("/:id", rooms.EditRoom)
		roomsGroup.DELETE("/:id", rooms.DeleteRoom)
	}
}

func indexPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}
