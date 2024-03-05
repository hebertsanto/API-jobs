package handlers

import (
	"github.com/gin-gonic/gin"
)

func Users(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List users",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "This route create a user",
	})
}

func ListUserById(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "This route list user by id",
	})
}
