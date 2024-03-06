package models

type UserProfile struct {
	UserName  string `json:"name"`
	Email     string `json:"email" validate:"required,email"`
	Position  string `json:"position" validate:"required"`
	GithubUrl string `json:"github_url"`
	Linkedin  string `json:"linkedin"`
	Website   string `json:"website"`
	UserId    string `json:"user_id" validate:"required"`
}
