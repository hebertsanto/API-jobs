package services

import (
	"vagas/infra/errors"
	repository "vagas/infra/repository/users"
	"vagas/models"
	"vagas/pkg/logger"

	"github.com/go-playground/validator"
)

type CreateUserService struct {
	Repo *repository.UserRepository
}

func (u *CreateUserService) CreateUser(user models.User) (models.User, error) {

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		logger.Log.Error("Error validating user...", err)
		return models.User{}, &errors.AppError{
			Code:    400,
			Message: "error validating user: " + err.Error(),
		}

	}

	err := u.Repo.CreateTableUsersIfNotExist()
	if err != nil {
		logger.Log.Errorf("Error creating user table: %v", err)
		return models.User{}, &errors.AppError{
			Code:    500,
			Message: "error creating user table" + err.Error(),
		}
	}

	user, err = u.Repo.CreateUser(user)

	if err != nil {
		logger.Log.Errorf("Error creating user: %v", err)
		return models.User{}, &errors.AppError{
			Code:    500,
			Message: "Error creating user: " + err.Error(),
		}
	}

	return user, nil
}

func GetUsers(u *CreateUserService) ([]models.User, error) {

	result, err := u.Repo.GetUsers()

	if err != nil {
		return nil, err
	}

	return result, nil

}
