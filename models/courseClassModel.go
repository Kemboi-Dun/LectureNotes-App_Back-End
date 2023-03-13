package models

import "gorm.io/gorm"

type CourseClass struct {
	gorm.Model
	Year     string
	Semester string
}
