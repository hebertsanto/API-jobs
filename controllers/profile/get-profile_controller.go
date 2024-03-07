package controllers

import (
	"vagas/database"
	"vagas/utils"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	db := database.GetDB()

	id := c.Param("id")

	if !utils.VerifyExistenceInDatabase(id, "profile") {
		c.JSON(404, gin.H{"error": "Profile_id not found in database"})
		return
	}

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
