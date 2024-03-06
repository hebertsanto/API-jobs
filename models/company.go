package models

type Company struct {
	ID             string `json:"id"`
	Name           string `json:"name" validate:"required"`
	Owner          string `json:"owner" validate:"required"`
	Cnpj           string `json:"cnpj" validate:"required"`
	TotalEmployees int    `json:"total_employees" validate:"required"`
	OpenVacancies  int    `json:"open_vacancies" validate:"required"`
	CreateAt       string `json:"create_at"`
}
