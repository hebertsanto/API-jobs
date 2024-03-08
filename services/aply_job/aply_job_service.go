package services

import (
	"vagas/infra/errors"
	repository "vagas/infra/repository/aply_job"
	"vagas/models"
	"vagas/pkg/logger"
)

type CreateAplyJobService struct {
	Repo *repository.AplyJobRepository
}

func (j *CreateAplyJobService) CreateAplyJobService(aply models.ApplyJob) (models.ApplyJob, error) {
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
