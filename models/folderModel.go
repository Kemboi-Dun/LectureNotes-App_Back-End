package models

import "gorm.io/gorm"

type Folder struct {
	gorm.Model
	Name           string
	ParentFolderID int
}
