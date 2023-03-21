package models

import "gorm.io/gorm"

type Documet struct {
	gorm.Model
	Name       string
	Path       string
	AuthorID   int
	FolderID   int
	Size       int64
	Year       int
	CourseType string
	CourseName string
	Semester   string
	UnitName   string
	UnitCode   string
	AuthorName string
	Type       string
}
