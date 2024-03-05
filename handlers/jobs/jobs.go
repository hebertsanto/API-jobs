package handlers

import (
	"github.com/gin-gonic/gin"
)

func ListJob(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List jobs",
	})
}

func PublicJob(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "This route Public a job",
	})
}
