package main

import (
	"auth_api/database"
	"auth_api/handlers"
	"auth_api/migration"
	"auth_api/repositories"
	"auth_api/routes"
	"auth_api/services"
	"log"
)

func main() {
	// Initialize database
	database.ConnectDatabase()
	migration.Migrate()

	// Initialize repositories
	userRepo := repositories.NewUserRepository()

	// Initialize services
	emailService := services.NewEmailService()
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo, emailService)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService, emailService)
	userHandler := handlers.NewUserHandler(userService)

	// Setup router
	r := routes.SetupRouter(authHandler, userHandler)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
