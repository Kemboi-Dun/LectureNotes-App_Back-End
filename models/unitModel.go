package models

import "gorm.io/gorm"

type Unit struct {
	gorm.Model

	Name       string
	Code       string
	Course     string
	CourseType string
}
