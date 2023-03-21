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
	initializers.DB.AutoMigrate(
		&models.User{},
		&models.Documet{},
		&models.FileAccess{},
		&models.Comment{},
		&models.Article{},
		&models.Doc{},
	)
}
