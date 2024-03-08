package services

import (
	"database/sql"
	"vagas/infra/errors"
	repository "vagas/infra/repository/jobs"
	"vagas/models"
	"vagas/pkg/logger"
)

type JobService struct {
	Repo *repository.JobRepository
}

func (j *JobService) CreateJobService(job models.Jobs) (models.Jobs, error) {
	err := j.Repo.CreateTableAplyIfNotExist()
	if err != nil {
		logger.Log.Error("Error creating job table...", err)
		return models.Jobs{}, &errors.AppError{
			Code:    500,
			Message: "error creating job table" + err.Error(),
		}
	}

	result, err := j.Repo.CreateJob(job)
	if err != nil {
		logger.Log.Error("Error creating job...", err)
		return models.Jobs{}, &errors.AppError{
			Code:    500,
			Message: "error creating job table" + err.Error(),
		}
	}

	return result, nil
}

func (j *JobService) GetJob(id string) (sql.Result, error) {

	result, err := j.Repo.GetJob(id)

	if err != nil {
		logger.Log.Errorf("Error getting user: %v", err)
		return result, &errors.AppError{
			Code:    500,
			Message: "Error getting job by id: " + err.Error(),
		}
	}

	return result, nil
}
