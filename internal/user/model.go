package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"index"`
	Password string
	Name     string
}
