package controller

import (
	"net/http"
	"vagas/infra/errors"
	repository "vagas/infra/repository/apply_job"
	"vagas/pkg/logger"
	services "vagas/services/apply_job"

	"github.com/gin-gonic/gin"
)

func DeleteAplyJob(c *gin.Context) {
	id := c.Param("id")

	aplyRepo := repository.NewAplyJobRepository()

	deleteAplyService := services.ApplyJob{Repo: aplyRepo}

	if err := deleteAplyService.DeleteAply(id); err != nil {
		logger.Log.Error("Error deleting aply job...", err)
		errors.HandlerError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "aply job deleted sucessfully",
	})
}
