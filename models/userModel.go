package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Role     string
	School   string
	Email    string
	Password string
}
