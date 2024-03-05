package handlers

import (
	"vagas/database"
	"vagas/models"

	"github.com/gin-gonic/gin"
)

func UpdateCompany(c *gin.Context) {

	db := database.GetDB()

	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not set"})
		return
	}

	var company models.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE company SET name = ?, owner = ?, cnpj = ?, total_employees = ?, open_vacancies = ?) VALUES ($1, $2, $3, $4, $5)"

	result, err := db.Exec(query, company.Name, company.Owner, company.Cnpj, company.TotalEmployees, company.OpenVacancies)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not update company" + err.Error()})
	}

	c.JSON(200, gin.H{
		"message": "company has been published in database",
		"result":  result,
	})
}
