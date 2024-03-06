package models

type Company struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Owner          string `json:"owner"`
	Cnpj           string `json:"cnpj"`
	TotalEmployees int    `json:"total_employees"`
	OpenVacancies  int    `json:"open_vacancies"`
	CreateAt       string `json:"create_at"`
}
