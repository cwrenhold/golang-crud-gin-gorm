package main

import (
	"golang-crud-gin-gorm/initializers"
	"golang-crud-gin-gorm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
