package dtos

type ResponseUserDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Cpf      string `json:"cpf"`
}

func NewResponseUserDTO(name string, password string, email string, cpf string) *ResponseUserDTO {
	return &ResponseUserDTO{
		Name:     name,
		Password: password,
		Email:    email,
		Cpf:      cpf,
	}
}
