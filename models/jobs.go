package models

type Jobs struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Salary      string `json:"salary"`
	Remote      bool   `json:"remote"`
	CompanyId   string `json:"company_id"`
	PublishedAt string `json:"published_at"`
	UpdateAt    string `json:"update_at"`
}
