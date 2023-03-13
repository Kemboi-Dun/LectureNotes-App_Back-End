package main

import (
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{},
		&models.Folder{},
		&models.File{},
		&models.FileAccess{},
		&models.Course{},
		&models.Year{},
		&models.CourseType{},
		&models.SemesterName{},
		&models.Semester{},
		&models.CourseClass{},
		&models.Unit{},
		&models.Comment{},
	)
}
