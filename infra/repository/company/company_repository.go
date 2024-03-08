package repository

import (
	"database/sql"
	"vagas/database"
	"vagas/models"
)

type CompanyRepostitory struct {
	DB *sql.DB
}

func NewCompanyRepository() *CompanyRepostitory {
	return &CompanyRepostitory{
		DB: database.GetDB(),
	}
}

func (c *CompanyRepostitory) CreateTableCompaniesIfNotExist() error {

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

func (u *CompanyRepostitory) CreateCompany(company models.Company) (models.Company, error) {
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
	err := u.DB.QueryRow(
		query,
		company.Name,
		company.Owner,
		company.Cnpj,
		company.TotalEmployees,
		company.OpenVacancies,
	).Scan(company.ID)

	if err != nil {
		return models.Company{}, err
	}

	return company, nil
}

func (u *CompanyRepostitory) GetCompanyById(id string) (models.Company, error) {
	var company models.Company
	query := `
	SELECT * FROM company WHERE id = ?
	`
	err := u.DB.QueryRow(query, id)

	if err != nil {
		return company, nil
	}

	return company, nil
}

func (u *CompanyRepostitory) DeleteCompany(id int) error {
	query := `
	DELETE FROM company WHERE id = ?
	`
	_, err := u.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (u *CompanyRepostitory) UpdateCompany(company models.Company, id string) (models.Company, error) {
	query := `
	  UPDATE company SET 
	     name = ?,
		 owner = ?, 
		 cnpj = ?,
		 total_employees = ?,
		 open_vacancies = ?
		 ) VALUES ($1, $2, $3, $4, $5)`

	err := u.DB.QueryRow(
		query,
		company.Name,
		company.Owner,
		company.Cnpj,
		company.TotalEmployees,
		company.OpenVacancies,
	).Scan(company.ID)
	if err != nil {
		return company, nil
	}
	return company, nil
}
