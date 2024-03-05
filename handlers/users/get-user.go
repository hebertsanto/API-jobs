package handlers

import "github.com/gin-gonic/gin"

func GetUserById(c *gin.Context) {

	id := c.Param("id")

	c.JSON(200, gin.H{
		"message": "this fun get user in database",
		"id":      id,
	})
}
