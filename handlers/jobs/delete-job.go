package handlers

import (
	"vagas/database"

	"github.com/gin-gonic/gin"
)

func DeleteJob(c *gin.Context) {

	db := database.GetDB()
	id := c.Param("id")

	query := "DELETE FROM jobs WHERE id = ?"

	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not set"})
		return
	}

	result, err := db.Exec(query, id)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not Deletin job" + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "job has been deleted",
		"result":  result,
	})
}
