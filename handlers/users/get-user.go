package handlers

import (
	"vagas/database"
	"vagas/repository"
	"vagas/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db := database.GetDB()

	if db == nil {
		c.JSON(500, gin.H{"error": "Error connecting to database"})
		return
	}
	userRepository := repository.NewUserRepository()
	userService := &services.NewUserService{Repo: userRepository}

	users, err := userService.Repo.GetUsers()

	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting users: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "User found successfully",
		"users":   users,
	})
}
