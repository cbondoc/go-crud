package main

import (
	"github.com/cbondoc/go-crud/initializers"
	"github.com/cbondoc/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Post{})
}
