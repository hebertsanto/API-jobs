package controllers

import (
	"net/http"
	"vagas/database"
	"vagas/infra/repository"
	"vagas/models"
	"vagas/pkg/logger"
	"vagas/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func CreateUser(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error connecting to database"})
		return
	}

	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Log.Infof("Payload received in invalid. Payload: %+v\n", user)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Payload received in invalid. Payload:" + err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		logger.Log.Infof("Error validating user: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "some error ocurred validating data" + err.Error(),
		})
		return
	}

	userRepository := repository.NewUserRepository()
	err := userRepository.CreateTableUsersIfNotExist()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user table " + err.Error()})
		return
	}

	userService := &services.CreateUserService{Repo: userRepository}

	user, err = userService.CreateUser(user)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user has been created": user,
	})
}
