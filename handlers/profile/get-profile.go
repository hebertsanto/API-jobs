package handlers

import (
	"vagas/database"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	db := database.GetDB()

	id := c.Param("id")

	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not set"})
		return
	}

	query := "SELECT * FROM profile WHERE id = ?"

	result, err := db.Exec(query, id)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not get profile" + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "profile found",
		"result":  result,
	})

}
