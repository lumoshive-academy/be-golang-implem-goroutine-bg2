package entity

type User struct {
	Model
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Photo    string `json:"photo" validate:"required"`
}
