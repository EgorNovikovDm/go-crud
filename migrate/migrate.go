package main

import (
	"github.com/EgoriyNovikovDm/go-crud/initializers"
	"github.com/EgoriyNovikovDm/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
