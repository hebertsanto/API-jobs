package controllers

import (
	"net/http"
	"vagas/infra/errors"
	repository "vagas/infra/repository/jobs"
	"vagas/models"
	"vagas/pkg/logger"
	services "vagas/services/jobs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func CreateJob(c *gin.Context) {
	job := models.Jobs{}

	if err := c.ShouldBindJSON(&job); err != nil {
		logger.Log.Error("Error binding job...", err)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()

	if err := validate.Struct(job); err != nil {
		logger.Log.Error("Error validating job...", err)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	jobRepo := repository.NewJobRepository()
	jobService := services.JobService{Repo: jobRepo}

	result, err := jobService.CreateJobService(job)

	if err != nil {
		logger.Log.Error("Error creating job...", err)
		errors.HandlerError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "job has been created",
		"result":  result,
	})

}
