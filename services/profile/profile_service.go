package services

import (
	"database/sql"
	"vagas/infra/errors"
	repository "vagas/infra/repository/profile"
	"vagas/models"
	"vagas/pkg/logger"

	"github.com/go-playground/validator"
)

type ProfileService struct {
	Repo *repository.ProfileRepository
}

func (p *ProfileService) CreateProfile(profile models.UserProfile) (models.UserProfile, error) {

	validate := validator.New()
	if err := validate.Struct(profile); err != nil {
		logger.Log.Error("Error validating profile...", err)
		return models.UserProfile{}, &errors.AppError{
			Code:    400,
			Message: "error validating profile: " + err.Error(),
		}

	}

	err := p.Repo.CreateTableProfileIfNotExist()
	if err != nil {
		logger.Log.Errorf("Error creating user table: %v", err)
		return models.UserProfile{}, &errors.AppError{
			Code:    500,
			Message: "error creating user table" + err.Error(),
		}
	}

	profile, err = p.Repo.CreateProfile(profile)

	if err != nil {
		logger.Log.Errorf("Error creating user: %v", err)
		return models.UserProfile{}, &errors.AppError{
			Code:    500,
			Message: "Error creating user: " + err.Error(),
		}
	}

	return profile, nil
}

func (p *ProfileService) GetProfile(id string) (sql.Result, error) {

	result, err := p.Repo.GetProfile(id)

	if err != nil {
		logger.Log.Errorf("Error getting user: %v", err)
		return result, &errors.AppError{
			Code:    500,
			Message: "Error getting user: " + err.Error(),
		}
	}

	return result, nil
}
