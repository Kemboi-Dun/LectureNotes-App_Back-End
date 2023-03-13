package models

import "gorm.io/gorm"

type FileAccess struct {
	gorm.Model
	FileID      int
	UserID      int
	AccessLevel string
}
