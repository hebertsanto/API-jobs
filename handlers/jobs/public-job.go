package handlers

import (
	"github.com/gin-gonic/gin"
)

func PublicJob(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "This route Public a job",
	})
}
