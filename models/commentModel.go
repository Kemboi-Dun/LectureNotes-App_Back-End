package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	AuthorName string
	Body       string
}
