package main

import (
	"github.com/EduardoMark/login-system-go/internal/config"
	"github.com/EduardoMark/login-system-go/internal/database"
	"github.com/EduardoMark/login-system-go/internal/models"
	"github.com/EduardoMark/login-system-go/internal/router"
)

func main() {
	// Initialize enviroment
	config.LoadEnv()

	// Setup DB
	database.SetupDatabase()
	database.AutoMigrate(models.User{})

	// Init routes
	router.SetupRoutes()
}
