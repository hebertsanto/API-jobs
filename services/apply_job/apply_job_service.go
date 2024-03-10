package services

import (
	"database/sql"
	"vagas/infra/errors"
	repository "vagas/infra/repository/apply_job"
	"vagas/models"
	"vagas/pkg/logger"
)

type ApplyJob struct {
	Repo *repository.AplyJobRepository
}

func (j *ApplyJob) CreateAplyJobService(aply models.ApplyJob) (models.ApplyJob, error) {

	err := j.Repo.CreateTableAplyIfNotExist()
	if err != nil {
		logger.Log.Error("Error creating aply job table...", err)
		return models.ApplyJob{}, &errors.AppError{
			Code:    500,
			Message: "error creating aply job table" + err.Error(),
		}
	}

	result, err := j.Repo.UserAplyJob(aply)
	if err != nil {
		logger.Log.Error("Error creating aply job...", err)
		return models.ApplyJob{}, &errors.AppError{
			Code:    500,
			Message: "error creating aply job table" + err.Error(),
		}
	}

	return result, nil
}

func (j *ApplyJob) DeleteAply(id string) error {
	err := j.Repo.UserDeleteAply(id)
	if err != nil {
		logger.Log.Error("Error deleting aply job...", err)
		return &errors.AppError{
			Code:    500,
			Message: "error deleting aply job table: " + err.Error(),
		}
	}

	return err
}

func (j *ApplyJob) GetAplyJob(id string) (sql.Result, error) {
	result, err := j.Repo.GetAplyJob(id)
	if err != nil {
		logger.Log.Error("Error deleting aply job...", err)
		return result, &errors.AppError{
			Code:    500,
			Message: "error deleting aply job table: " + err.Error(),
		}
	}

	return result, err
}
