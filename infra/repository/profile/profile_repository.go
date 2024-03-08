package repository

import (
	"database/sql"
	"vagas/database"
	"vagas/models"
	"vagas/pkg/logger"
)

type ProfileRepository struct {
	db *sql.DB
}

func CreateProfileRepository() *ProfileRepository {
	return &ProfileRepository{
		db: database.GetDB(),
	}
}

func (p *ProfileRepository) CreateTableProfileIfNotExist() error {
	createTableQuerySQL := `
	CREATE TABLE IF NOT EXISTS profiles (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL
		position VARCHAR(255) NOT NULL,
		github_url VARCHAR(255) NOT NULL,
		linkedin VARCHAR(255) NOT NULL,
		website VARCHAR(255) NOT NULL,
		user_id VARCHAR(255) NOT NULL
		FOREING KEY (user_id) REFERENCES users(id)
	)
	`
	_, err := p.db.Exec(createTableQuerySQL)
	if err != nil {
		logger.Log.Error("Error creating table profiles...", err)
		return err
	}
	return nil
}

func (p *ProfileRepository) CreateProfile(profile models.UserProfile) (models.UserProfile, error) {
	query := `INSERT INTO profiles (name, email, position, github_url, linkedin, website, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING name`
	var name string
	err := p.db.QueryRow(query, profile.UserName, profile.Email, profile.Position, profile.GithubUrl, profile.Linkedin, profile.Website, profile.UserId).Scan(&name)
	if err != nil {
		logger.Log.Error("Error creating profile...", err)
		return models.UserProfile{}, err
	}

	profile.UserName = name

	return profile, nil

}

func (p *ProfileRepository) GetProfile(profile models.UserProfile, id string) (sql.Result, error) {

	query := `SELECT * FROM profiles WHERE id = ?`

	result, err := p.db.Exec(query, id)
	if err != nil {
		logger.Log.Error("Error getting profile...", err)
		return result, err
	}

	return result, nil

}
