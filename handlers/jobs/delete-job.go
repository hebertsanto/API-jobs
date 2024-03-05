package handlers

import "github.com/gin-gonic/gin"

func DeleteJob(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this func delete job in database",
	})
}
