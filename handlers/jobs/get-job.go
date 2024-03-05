package handlers

import "github.com/gin-gonic/gin"

func GetJobById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this fun get job in database",
	})
}
