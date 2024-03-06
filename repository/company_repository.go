package repository

import (
	"database/sql"
	"log"
	"vagas/database"
	"vagas/models"
)

type NewCompanyRepository struct {
	DB *sql.DB
}

func (c *NewCompanyRepository) CreateTableCompaniesIfNotExist() error {
	db := database.GetDB()

	if db == nil {
		log.Fatal("Database not found")
	}

	query := `
	CREATE TABLE IF NOT EXISTS company (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		owner VARCHAR(100) NOT NULL,
		cnpj VARCHAR(14) NOT NULL,
		total_employees INT NOT NULL,
		open_vacancies INT NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := c.DB.Exec(query)

	if err != nil {
		return err
	}
	return nil
}

func (u *NewCompanyRepository) CreateUser(company models.Company) error {

	db := database.GetDB()

	if db == nil {
		log.Fatal("Database not found")
	}

	query := `
	INSERT INTO company (
		name, 
		owner, 
		cnpj, 
		total_employees, 
		open_vacancies
	  ) 
	  VALUES ($1, $2, $3, $4, $5)
	`

	var err error

	_, err = u.DB.Exec(
		query,
		company.Name,
		company.Owner,
		company.Cnpj,
		company.TotalEmployees,
		company.OpenVacancies,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *NewCompanyRepository) GetCompanyById(id int) models.Company {
	var company models.Company
	query := `
	SELECT * FROM company WHERE id = ?
	`
	err := u.DB.QueryRow(query, id)

	if err != nil {
		return company
	}

	return company
}
