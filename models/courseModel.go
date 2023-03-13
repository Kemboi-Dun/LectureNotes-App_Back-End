package models

import "gorm.io/gorm"

// File model definition
type Course struct {
	gorm.Model
	Name string
}
