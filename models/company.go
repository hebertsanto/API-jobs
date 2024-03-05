package models

type Company struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Owner          string `json:"owner"`
	TotalEmployees int    `json:"total_employees"`
	OpenVacancies  int    `json:"open_vacancies"`
	CreateAt       string `json:"create_at"`
}
