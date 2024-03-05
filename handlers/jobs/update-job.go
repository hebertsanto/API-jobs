package handlers

import "github.com/gin-gonic/gin"

func UpdateJobById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this func update job in database",
	})
}
