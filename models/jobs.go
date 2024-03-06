package models

type Jobs struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Company     string `json:"company" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Salary      string `json:"salary" validate:"required"`
	Remote      bool   `json:"remote" validate:"required"`
	CompanyId   string `json:"company_id" validate:"required"`
	PublishedAt string `json:"published_at"`
	UpdateAt    string `json:"update_at"`
}
