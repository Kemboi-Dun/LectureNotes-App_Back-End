package models

import "gorm.io/gorm"

type Semester struct {
	gorm.Model
	Course       string
	CourseType   string
	SemesterName string
}
