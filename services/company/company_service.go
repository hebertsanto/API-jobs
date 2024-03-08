package services

import (
	"vagas/infra/errors"
	repository "vagas/infra/repository/company"
	"vagas/models"
	"vagas/pkg/logger"
)

type CompanyService struct {
	Repo *repository.CompanyRepostitory
}

func (c *CompanyService) CreateCompany(company models.Company) (models.Company, error) {

	err := c.Repo.CreateTableCompaniesIfNotExist()

	if err != nil {
		logger.Log.Errorf("Error creating company table: %v", err)
		return models.Company{}, &errors.AppError{
			Code:    500,
			Message: "error creating company table" + err.Error(),
		}
	}

	result, err := c.Repo.CreateCompany(company)
	if err != nil {
		logger.Log.Errorf("Error creating company: %v", err)
		return models.Company{}, &errors.AppError{
			Code:    500,
			Message: "Error creating company: " + err.Error(),
		}
	}

	logger.Log.Infof("Company created: %v", result)
	return result, nil
}

func (c *CompanyService) DeleteCompany(id int) error {
	err := c.Repo.DeleteCompany(id)
	if err != nil {
		logger.Log.Errorf("Error deleting company: %v", err)
		return &errors.AppError{
			Code:    500,
			Message: "Error deleting company: " + err.Error(),
		}
	}
	return nil
}
