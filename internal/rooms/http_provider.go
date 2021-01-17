package rooms

import (
	"github.com/gin-gonic/gin"
)

// GetRooms ...
func GetRooms(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Room list",
	})
}

// EditRoom ...
func EditRoom(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Room edit",
	})
}

// CreateRoom ...
func CreateRoom(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Room create",
	})
}

// DeleteRoom ...
func DeleteRoom(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Room delete",
	})
}

// InviteUser ...
func InviteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User invite to room",
	})
}

// RemoveUser ...
func RemoveUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User remove from room",
	})
}
