package utils

import (
	"fmt"
	"vagas/database"
)

func VerifyExistenceInDatabase(id string, table string) bool {
	db := database.GetDB()
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE id = ?", table)

	var count int
	err := db.QueryRow(query, id).Scan(&count)

	if err != nil {
		return false
	}

	return count > 0
}
