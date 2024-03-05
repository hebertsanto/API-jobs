package handlers

import (
	"database/sql"
	"net/http"
	"vagas/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *sql.DB

func CreateUser(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Name == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Email are required"})
		return
	}

	c.JSON(201, gin.H{
		"message": "user create sucessfully",
		"user":    user,
	})
}

func Users(c *gin.Context) {

	_, err := db.Exec("SELECT * FROM users")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "some error ocurred ",
		})
	}
	c.JSON(200, gin.H{
		"message": "List users",
	})
}

func ListUserById(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "This route list user by id",
	})
}
