package controllers

import (
	"net/http"
	"strconv"
	"vagas/infra/errors"
	repository "vagas/infra/repository/company"
	"vagas/pkg/logger"
	services "vagas/services/company"

	"github.com/gin-gonic/gin"
)

func DeleteCompany(c *gin.Context) {

	id := c.Param("id")

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
