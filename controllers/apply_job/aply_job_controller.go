package controller

import (
	"net/http"
	"vagas/infra/errors"
	repository "vagas/infra/repository/apply_job"
	"vagas/models"
	"vagas/pkg/logger"
	services "vagas/services/apply_job"

	"github.com/gin-gonic/gin"
)

func CreateAplyJob(c *gin.Context) {
	aply := models.ApplyJob{}

	if err := c.BindJSON(&aply); err != nil {
		logger.Log.Error("Error binding json...", err)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	aplyRepo := repository.NewAplyJobRepository()

	aplyService := services.ApplyJob{Repo: aplyRepo}

	aply, err := aplyService.CreateAplyJobService(aply)

	if err != nil {
		logger.Log.Error("Error creating aply job...", err)
		errors.HandlerError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "aply job created sucessfully",
		"result":  aply,
	})
}
