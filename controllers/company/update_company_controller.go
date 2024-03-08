package controllers

import (
	"net/http"
	"vagas/infra/errors"
	repository "vagas/infra/repository/company"
	"vagas/models"
	"vagas/pkg/logger"
	services "vagas/services/company"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func UpdateCompany(c *gin.Context) {
	company := models.Company{}

	id := c.Param("id")
	if err := c.ShouldBindJSON(&company); err != nil {
		logger.Log.Errorf("Error binding company: %v", err)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()

	if err := validate.Struct(company); err != nil {
		logger.Log.Errorf("Error validating company: %v", err)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	companyRepo := repository.NewCompanyRepository()
	companyService := services.CompanyService{Repo: companyRepo}

	result, err := companyService.UpdateCompany(company, id)
	if err != nil {
		logger.Log.Errorf("Error updating company: %v", err)
		errors.HandlerError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Company updated successfully",
		"company": result,
	})

}
