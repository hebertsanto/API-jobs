package handlers

import (
	"database/sql"
	"vagas/database"
	"vagas/models"

	"github.com/gin-gonic/gin"
)

func CreateTableQuerySql(db *sql.DB) error {
	createTableQuerySQL := `CREATE TABLE IF NOT EXISTS profile (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		position VARCHAR(255) NOT NULL,
		github_url VARCHAR(255) NOT NULL,
		linkedin_url VARCHAR(255) NOT NULL,
		website_url VARCHAR(255) NOT NULL,
		user_id VARCHAR(255) NOT NULL,
		FOREING KEY (user_id) REFERENCES users (id)

	)`
	_, err := db.Exec(createTableQuerySQL)

	if err != nil {
		return err
	}

	return nil
}

func CreateProfile(c *gin.Context) {
	db := database.GetDB()

	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not set"})
	}

	err := CreateTableQuerySql(db)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not create profile" + err.Error()})
		return
	}

	var profile models.UserProfile

	query := "INSERT INTO profile (username, email, position, github_url, linkedin_url, website_url, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	result, err := db.Exec(query, profile.UserName, profile.Email, profile.Position, profile.GithubUrl, profile.Linkedin, profile.Website, profile.UserId)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not create profile" + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Profile created",
		"result":  result,
	})
}
