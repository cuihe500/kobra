package dto

type UserRegisterDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name,omitempty"`
	Age      uint   `json:"age,omitempty"`
}
