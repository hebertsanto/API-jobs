package repository

import (
	"database/sql"
	"vagas/models"
)

type NewUserRepository struct {
	DB *sql.DB
}

func (u *NewUserRepository) CreateTableUsersIfNotExist() error {
	createTableQuerySQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		cpf VARCHAR(255) NOT NULL
	)
	`
	_, err := u.DB.Exec(createTableQuerySQL)
	if err != nil {
		return err
	}
	return nil
}

func (u *NewUserRepository) CreateUser(user models.User) (sql.Result, error) {

	query := `INSERT INTO users (name, password, email, cpf ) VALUES ($1, $2, $3, $4)`

	result, err := u.DB.Exec(
		query,
		user.Name,
		user.Password,
		user.Email,
		user.Cpf,
	)

	if err != nil {
		return result, err
	}

	return result, err
}
