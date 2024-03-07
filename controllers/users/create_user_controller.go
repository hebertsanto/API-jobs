package controllers

import (
	"net/http"
	"vagas/infra/errors"
	repository "vagas/infra/repository/users"
	"vagas/models"
	"vagas/pkg/logger"
	services "vagas/services/users"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func CreateUser(c *gin.Context) {

	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Log.Infof("Payload received in invalid. Payload: %+v\n", user)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		logger.Log.Infof("error validating user data: %+v\n", user)
		errors.HandlerError(c, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	userRepository := repository.NewUserRepository()

	userService := services.CreateUserService{Repo: userRepository}

	user, err := userService.CreateUser(user)

	if err != nil {
		logger.Log.Error("Error creating user...", err)
		errors.HandlerError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user has been created": user,
	})
}
