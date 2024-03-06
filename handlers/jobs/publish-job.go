package handlers

import (
	"database/sql"
	"net/http"
	"vagas/database"
	"vagas/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateTableQuerySql(db *sql.DB) error {
	createTableQuerySQL := `
		CREATE TABLE IF NOT EXISTS jobs (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			company VARCHAR(255) NOT NULL,
			location VARCHAR(255) NOT NULL,
			salary REAL NOT NULL,
			remote BOOLEAN NOT NULL
			comapany_id VARCHAR(255) NOT NULL,
			published_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			FOREIGN KEY (company_id) REFERENCES company (id)
		)
	`
	_, err := db.Exec(createTableQuerySQL)
	if err != nil {
		return err
	}

	return nil

}

func PublishJob(c *gin.Context) {
	db := database.GetDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not set"})
		return
	}

	err := CreateTableQuerySql(db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating jobs table: " + err.Error()})
		return
	}

	var job models.Jobs

	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(job); err != nil {
		c.JSON(400, gin.H{"error": "error validatig data" + err.Error()})
		return
	}

	query := `
	INSERT INTO jobs (
	 name, 
	 description, 
	 company, 
	 location, 
	 salary, 
	 remote, 
	 company_id
	 ) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	result, err := db.Exec(
		query,
		job.Name,
		job.Description,
		job.Company,
		job.Location,
		job.Salary,
		job.Remote,
		job.CompanyId,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting job: " + err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "This route Public a job",
		"result":  result,
	})
}
