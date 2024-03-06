package controllers

import (
	"vagas/database"
	"vagas/utils"

	"github.com/gin-gonic/gin"
)

func DeleteJob(c *gin.Context) {

	db := database.GetDB()
	id := c.Param("id")

	if !utils.VerifyExistenceInDatabase(id, "jobs") {
		c.JSON(404, gin.H{"error": "Job_id not found in database"})
		return
	}
	query := "DELETE FROM jobs WHERE id = ?"

	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not set"})
		return
	}

	var err error
	_, err = db.Exec(query, id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not delete company: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "job has been deleted",
	})
}
