package handlers

import (
	"database/sql"
	"net/http"
	"vagas/database"
	"vagas/models"

	"github.com/gin-gonic/gin"
)

func ensureUsersTableExists(db *sql.DB) error {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL
		)
	`
	_, err := db.Exec(createTableSQL)
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

	err := ensureUsersTableExists(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating users table: " + err.Error()})
		return
	}

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Name == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Email are required"})
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
