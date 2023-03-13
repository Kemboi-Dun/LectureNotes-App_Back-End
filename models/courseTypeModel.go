package models

import "gorm.io/gorm"

type CourseType struct {
	gorm.Model
	Name string
}
