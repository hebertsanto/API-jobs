package repository

import (
	"database/sql"
	"vagas/database"
	"vagas/models"
	"vagas/pkg/logger"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.GetDB(),
	}
}

func (u *UserRepository) CreateTableUsersIfNotExist() error {

	createTableQuerySQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		cpf VARCHAR(255) NOT NULL
	)
	`
	_, err := u.db.Exec(createTableQuerySQL)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) CreateUser(user models.User) (models.User, error) {

	query := `INSERT INTO usuarios (name, password, email, cpf ) VALUES ($1, $2, $3, $4) RETURNING name`
	var name string
	err := u.db.QueryRow(query, user.Name, user.Password, user.Email, user.Cpf).Scan(&name)
	if err != nil {
		return models.User{}, err
	}

	user.Name = name

	return user, nil

}

func (u *UserRepository) GetUsers() ([]models.User, error) {

	query := `SELECT name, password, email, cpf FROM usuarios`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Name, &user.Password, &user.Email, &user.Cpf); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		logger.Log.Error("Error on get users", err)
		return nil, err
	}

	return users, nil
}
