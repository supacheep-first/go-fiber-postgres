package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" validate:"required,min=2"`
	Email string `json:"email" validate:"required,email"`
}
