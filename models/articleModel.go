package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model

	AuthorID   int
	Tag        string
	Title      string
	AuthorName string
	Body       string
}
