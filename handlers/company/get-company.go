package handlers

import (
	"vagas/database"
	"vagas/utils"

	"github.com/gin-gonic/gin"
)

func GetCompany(c *gin.Context) {
	db := database.GetDB()

	id := c.Param("id")
	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not set"})
		return
	}

	if !utils.VerifyExistenceInDatabase(id, "profile") {
		c.JSON(404, gin.H{"error": "Profile_id not found in database"})
		return
	}

	query := "SELECT * FROM company WHERE id = ?"

	result, err := db.Exec(query, id)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not delete company" + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "company found",
		"result":  result,
	})
}
