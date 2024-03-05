package handlers

import (
	"vagas/database"
	"vagas/models"

	"github.com/gin-gonic/gin"
)

func UpdateJobById(c *gin.Context) {
	db := database.GetDB()

	var jobs models.Jobs

	query := "UPDATE jobs SET name = ?, description = ?, company = ?, location = ?, salary = ?, remote = ?,  company_id = ? WHERE id = ?"

	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not set"})
	}

	result, err := db.Exec(query, jobs.Name, jobs.Description, jobs.Company, jobs.Location, jobs.Salary, jobs.Remote, jobs.CompanyId)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not update job" + err.Error()})
	}

	c.JSON(200, gin.H{
		"message": "this func update job in database",
		"result":  result,
	})
}
