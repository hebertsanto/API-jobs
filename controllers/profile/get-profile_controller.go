package controllers

import (
	"vagas/infra/errors"
	repository "vagas/infra/repository/profile"
	"vagas/models"
	"vagas/pkg/logger"
	services "vagas/services/profile"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func GetProfile(c *gin.Context) {

	id := c.Param("id")
	profile := models.UserProfile{}

	if err := c.ShouldBindJSON(&profile); err != nil {
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), 400)
		return
	}

	validate := validator.New()

	if err := validate.Struct(profile); err != nil {
		logger.Log.Infof("error validating user data: %+v\n", profile)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), 400)
		return
	}

	profileRepo := repository.CreateProfileRepository()

	profileService := services.ProfileService{Repo: profileRepo}

	result, err := profileService.GetProfile(profile, id)

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
