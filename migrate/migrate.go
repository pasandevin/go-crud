package main

import (
	"github.com/pasandevin/go-crud/initializers"
	"github.com/pasandevin/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
