package handlers

import "github.com/gin-gonic/gin"

func GetUserById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this fun get user in database",
	})
}
