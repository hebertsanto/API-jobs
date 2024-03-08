package controllers

import (
	"net/http"
	"vagas/infra/errors"
	repository "vagas/infra/repository/jobs"
	"vagas/pkg/logger"

	"github.com/gin-gonic/gin"
)

func GetJobById(c *gin.Context) {
	id := c.Param("id")

	jobRepo := repository.NewJobRepository()

	result, err := jobRepo.GetJob(id)
	if err != nil {
		logger.Log.Error("Error getting job...", err)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "job found",
		"result":  result,
	})
}
