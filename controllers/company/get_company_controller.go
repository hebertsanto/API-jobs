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

func GetCompanyById(c *gin.Context) {

	company := models.Company{}
	if err := c.ShouldBindJSON(&company); err != nil {
		logger.Log.Errorf("Payload received in invalid. Payload: %+v\n", company)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(company); err != nil {
		logger.Log.Errorf("error validating company data: %+v\n", company)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
	}

	companyRepo := repository.NewCompanyRepository()
	userService := services.CompanyService{Repo: companyRepo}

	result, err := userService.GetCompanyService(company.ID)
	if err != nil {
		logger.Log.Errorf("Error getting company: %v", err)
		errors.HandlerError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"message": "company found",
		"company": result,
	})

}
