package controllers

import (
	"net/http"
	"vagas/database"
	"vagas/infra/errors"
	repository "vagas/infra/repository/users"
	"vagas/pkg/logger"
	services "vagas/services/users"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db := database.GetDB()

	if db == nil {
		c.JSON(500, gin.H{"error": "Error connecting to database"})
		return
	}
	userRepository := repository.NewUserRepository()
	userService := &services.CreateUserService{Repo: userRepository}

	users, err := userService.Repo.GetUsers()

	if err != nil {
		logger.Log.Error("Error getting All users...", err)
		errors.HandlerError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"message": "All users found successfully",
		"users":   users,
	})
}
