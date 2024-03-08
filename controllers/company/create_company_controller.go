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

func PulishCompany(c *gin.Context) {
	company := models.Company{}

	if err := c.ShouldBindJSON(&company); err != nil {
		logger.Log.Infof("Payload received in invalid. Payload: %+v\n", company)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	validator := validator.New()
	if err := validator.Struct(company); err != nil {
		logger.Log.Infof("error validating company data: %+v\n", company)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	userRepository := repository.NewCompanyRepository()

	userService := services.CompanyService{Repo: userRepository}

	company, err := userService.CreateCompany(company)

	if err != nil {
		logger.Log.Error("Error creating company...", err)
		errors.HandlerError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "company created sucessfully",
		"company": company,
	})
}
