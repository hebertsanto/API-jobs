package handlers

import (
	"database/sql"
	"net/http"
	"vagas/database"
	"vagas/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateTableUsersIfNotExist(db *sql.DB) error {
	createTableQuerySQL := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL
			cpf VARCHAR(255) NOT NULL
		)
	`
	_, err := db.Exec(createTableQuerySQL)
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(c *gin.Context) {

	db := database.GetDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not set"})
		return
	}

	err := CreateTableUsersIfNotExist(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating users table: " + err.Error()})
		return
	}

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		c.JSON(400, gin.H{"error": "some error ocurred validating data" + err.Error()})
		return
	}

	query := "INSERT INTO users (name, email) VALUES ($1, $2)"

	result, err := db.Exec(query, user.Name, user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting user into database" + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully in database",
		"user":    result,
	})
}
