package messages

import (
	"github.com/gin-gonic/gin"
)

// GetMessages ...
func GetMessages(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Messages",
	})
}

// SendMessage ...
func SendMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Send message",
	})
}

// EditMessage ...
func EditMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Send message",
	})
}

// DeleteMessage ...
func DeleteMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Message",
	})
}
