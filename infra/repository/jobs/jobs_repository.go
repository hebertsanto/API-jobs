package repository

import (
	"database/sql"
	"vagas/database"
	"vagas/models"
	"vagas/pkg/logger"
)

type JobRepository struct {
	DB *sql.DB
}

func NewJobRepository() *JobRepository {
	return &JobRepository{
		DB: database.GetDB(),
	}
}

func (j *JobRepository) CreateTableAplyIfNotExist() error {
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
	_, err := j.DB.Exec(createTableQuerySQL)
	if err != nil {
		logger.Log.Error("Error creating table jobs...", err)
		return err
	}
	return nil
}

func (j *JobRepository) CreateJob(job models.Jobs) (models.Jobs, error) {
	query := `
	INSERT INTO jobs (name, description, company, location, salary, remote, company_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`
	err := j.DB.QueryRow(query, job.Name, job.Description, job.Company, job.Location, job.Salary, job.Remote, job.ID).Scan(&job.ID)
	if err != nil {
		logger.Log.Error("Error creating job...", err)
		return models.Jobs{}, nil
	}
	return job, nil
}

func (j *JobRepository) GetJob(id string) (sql.Result, error) {

	query := `SELECT * FROM jobs WHERE id = ?`

	result, err := j.DB.Exec(query, id)
	if err != nil {
		logger.Log.Error("Error getting job...", err)
		return result, err
	}

	return result, nil

}
