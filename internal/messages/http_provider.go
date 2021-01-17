package messages

import (
	"github.com/gin-gonic/gin"
)

// HTTPProvider ...
type HTTPProvider struct {
	Engine *gin.Engine
}

// NewHTTPProvider ...
func NewHTTPProvider(e *gin.Engine) *HTTPProvider {
	h := HTTPProvider{Engine: e}

	group := h.Engine.Group("/messages")
	{
		group.GET("/", h.getMessages)
		group.POST("/", h.sendMessage)
		group.PUT("/:id", h.editMessage)
		group.DELETE("/:id", h.deleteMessage)
	}

	return &h
}

// GetMessages ...
func (hp *HTTPProvider) getMessages(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Messages",
	})
}

// SendMessage ...
func (hp *HTTPProvider) sendMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Send message",
	})
}

// EditMessage ...
func (hp *HTTPProvider) editMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Send message",
	})
}

// DeleteMessage ...
func (hp *HTTPProvider) deleteMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Message",
	})
}
