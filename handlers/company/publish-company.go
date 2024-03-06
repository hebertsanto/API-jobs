package handlers

import (
	"database/sql"
	"vagas/database"
	"vagas/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateTableCompanyQuerySQL(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS company (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		owner VARCHAR(100) NOT NULL,
		cnpj VARCHAR(14) NOT NULL,
		total_employees INT NOT NULL,
		open_vacancies INT NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func PulishCompany(c *gin.Context) {

	db := database.GetDB()

	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not set"})
		return
	}

	err := CreateTableCompanyQuerySQL(db)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating company table: " + err.Error()})
		return
	}

	var company models.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(company); err != nil {
		c.JSON(400, gin.H{"error": "some error ocurred validating data" + err.Error()})
		return
	}

	query := `
	INSERT INTO company (
		name, 
		owner, 
		cnpj, 
		total_employees, 
		open_vacancies
	  ) 
		VALUES ($1, $2, $3, $4, $5)`

	result, err := db.Exec(
		query,
		company.Name,
		company.Owner,
		company.Cnpj,
		company.TotalEmployees,
		company.OpenVacancies,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not create company" + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "company has been published in database",
		"result":  result,
	})
}
