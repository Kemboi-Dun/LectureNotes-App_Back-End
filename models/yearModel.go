package models

import "gorm.io/gorm"

type Year struct {
	gorm.Model
	Name string
}
