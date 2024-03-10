package controllers

import (
	"vagas/infra/errors"
	repository "vagas/infra/repository/profile"
	"vagas/models"
	"vagas/pkg/logger"
	services "vagas/services/profile"

	"github.com/gin-gonic/gin"
)

func CreateProfile(c *gin.Context) {

	profile := models.UserProfile{}
	if err := c.ShouldBindJSON(&profile); err != nil {
		logger.Log.Infof("Payload received in invalid. Payload: %+v\n", profile)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), 400)
		return
	}

	profileRepo := repository.CreateProfileRepository()

	userService := services.ProfileService{Repo: profileRepo}

	profile, err := userService.CreateProfile(profile)

	if err != nil {
		logger.Log.Error("Error creating profile...", err)
		errors.HandlerError(c, "INTERNAL_SERVER_ERROR", err.Error(), 500)
		return
	}

	c.JSON(201, gin.H{
		"message": "profile has been created",
		"profile": profile,
	})

}
