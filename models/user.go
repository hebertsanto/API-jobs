package models

type User struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Cpf      string `json:"cpf" validate:"required"`
}
