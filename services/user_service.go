package services

import (
	"vagas/models"
	"vagas/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NewUserService struct {
	Repo *repository.UserRepository
}

func (u *NewUserService) CreateUser(c *gin.Context, user models.User) {

	err := u.Repo.CreateTableUsersIfNotExist()

	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating user table: " + err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		c.JSON(400, gin.H{"error": "some error ocurred validating data" + err.Error()})
		return
	}

	user, err = u.Repo.CreateUser(user)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating user: " + err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func GetUsers(u *NewUserService) ([]models.User, error) {

	result, err := u.Repo.GetUsers()

	if err != nil {
		return nil, err
	}

	return result, nil

}
