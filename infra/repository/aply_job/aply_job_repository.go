package repository

import (
	"database/sql"
	"vagas/database"
	"vagas/models"
	"vagas/pkg/logger"
)

type AplyJobRepository struct {
	DB *sql.DB
}

func NewAplyJobRepository() *AplyJobRepository {
	return &AplyJobRepository{
		DB: database.GetDB(),
	}
}

func (a *AplyJobRepository) CreateTableAplyIfNotExist() error {
	createTableQuerySQL := `
	CREATE TABLE IF NOT EXISTS aply_job (
		id SERIAL PRIMARY KEY,
		userName VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		job_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (job_id) REFERENCES jobs(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`
	_, err := a.DB.Exec(createTableQuerySQL)
	if err != nil {
		return err
	}
	return nil
}

func (a *AplyJobRepository) UserAplyJob(aply models.ApplyJob) (models.ApplyJob, error) {
	query := `
	INSERT INTO aply_job (userName, email, job_id, user_id) VALUES ($1, $2, $3) RETURNING id
	`
	err := a.DB.QueryRow(query, aply.UserName, aply.Email, aply.JobID, aply.UserID)
	if err != nil {
		return models.ApplyJob{}, nil
	}
	return aply, nil
}

func (a *AplyJobRepository) UserDeleteAply(id string) error {

	query := `
	DELETE FROM aply_job WHERE id = ?
	`
	_, err := a.DB.Exec(query, id)

	return err
}

func (a *AplyJobRepository) GetAplyJob(id string) (sql.Result, error) {

	query := `
	SELECT FROM aply_job WHERE id = ?
	`
	result, err := a.DB.Exec(query, id)
	if err != nil {
		logger.Log.Error("Error getting aply job...", err)
		return nil, err
	}
	return result, err
}
