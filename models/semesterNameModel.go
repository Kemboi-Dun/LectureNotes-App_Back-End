package models

import "gorm.io/gorm"

type SemesterName struct {
	gorm.Model
	Name string
}
