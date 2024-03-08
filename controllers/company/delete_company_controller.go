package controllers

import (
	"net/http"
	"strconv"
	"vagas/infra/errors"
	repository "vagas/infra/repository/company"
	"vagas/models"
	"vagas/pkg/logger"
	services "vagas/services/company"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func DeleteCompany(c *gin.Context) {

	id := c.Param("id")

	company := models.Company{}
	if err := c.ShouldBindJSON(&company); err != nil {
		logger.Log.Infof("Payload received in invalid. Payload: %+v\n", company)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(company); err != nil {
		logger.Log.Infof("error validating company data: %+v\n", company)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
	}
	companyRepo := repository.NewCompanyRepository()

	companyService := services.CompanyService{Repo: companyRepo}

	idParam, err := strconv.Atoi(id)
	if err != nil {
		logger.Log.Errorf("Error converting id to int: %v", err)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	err = companyService.DeleteCompany(idParam)
	if err != nil {
		errors.HandlerError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"message": "company has been deleted",
	})
}
