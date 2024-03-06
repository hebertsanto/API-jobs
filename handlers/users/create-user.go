package handlers

import (
	"vagas/database"
	"vagas/infra/repository"
	"vagas/models"
	"vagas/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	db := database.GetDB()

	if db == nil {
		c.JSON(500, gin.H{"error": "Error connecting to database"})
		return

	}

	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "some error ocurred validating data" + err.Error()})
	}

	userRepository := repository.NewUserRepository()
	err := userRepository.CreateTableUsersIfNotExist()

	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating user table: " + err.Error()})
		return
	}

	userService := &services.NewUserService{Repo: userRepository}
	userService.CreateUser(c, user)
}
