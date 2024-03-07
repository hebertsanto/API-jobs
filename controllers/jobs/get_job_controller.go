package controllers

import (
	"vagas/database"

	"github.com/gin-gonic/gin"
)

func GetJobById(c *gin.Context) {

	db := database.GetDB()

	params := c.Param("id")

	query := "SELECT * FROM jobs WHERE id = ? "

	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not set"})
		return

	}

	job, err := db.Exec(query, params)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting job table: " + err.Error()})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting job: " + err.Error()})
	}
	c.JSON(200, gin.H{
		"message": "this fun get job in database",
		"job":     job,
	})
}
