package controllers

import (
	"vagas/infra/errors"
	repository "vagas/infra/repository/profile"
	"vagas/pkg/logger"
	services "vagas/services/profile"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {

	id := c.Param("id")
	profileRepo := repository.CreateProfileRepository()

	profileService := services.ProfileService{Repo: profileRepo}

	result, err := profileService.GetProfile(id)

	if err != nil {
		logger.Log.Error("Error getting profile...", err)
		errors.HandlerError(c, "INTERNAL_SERVER_ERROR", err.Error(), 500)
		return
	}

	c.JSON(200, gin.H{
		"message": "profile found",
		"result":  result,
	})

}
