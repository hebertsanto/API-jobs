package controller

import (
	"net/http"
	"vagas/infra/errors"
	repository "vagas/infra/repository/aply_job"
	"vagas/models"
	"vagas/pkg/logger"
	services "vagas/services/aply_job"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func DeleteAplyJob(c *gin.Context) {
	aply := models.ApplyJob{}
	id := c.Param("id")
	if err := c.ShouldBindJSON(&aply); err != nil {
		logger.Log.Error("Error binding json...", err)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()

	if err := validate.Struct(aply); err != nil {
		logger.Log.Error("Error validating aply job...", err)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	aplyRepo := repository.NewAplyJobRepository()

	deleteAplyService := services.AplyJobService{Repo: aplyRepo}

	if err := deleteAplyService.DeleteAply(id); err != nil {
		logger.Log.Error("Error deleting aply job...", err)
		errors.HandlerError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "aply job deleted sucessfully",
	})
}
