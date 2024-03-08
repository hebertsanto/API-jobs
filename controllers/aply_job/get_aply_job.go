package controller

import (
	"net/http"
	"vagas/infra/errors"
	repository "vagas/infra/repository/aply_job"
	"vagas/pkg/logger"

	"github.com/gin-gonic/gin"
)

func GetAplyJob(c *gin.Context) {
	id := c.Param("id")

	aplyRepo := repository.NewAplyJobRepository()
	result, err := aplyRepo.GetAplyJob(id)
	if err != nil {
		logger.Log.Error("Error getting aply job...", err)
		errors.HandlerError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "aply job retrieved sucessfully",
		"result":  result,
	})
}
