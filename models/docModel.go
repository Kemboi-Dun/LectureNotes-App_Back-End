package models

import "gorm.io/gorm"

type Doc struct {
	gorm.Model

	Path     string
	AuthorID int
}
