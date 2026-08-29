package main

import (
	"log"

	"ntl-test/backend/internal/config"
	"ntl-test/backend/internal/database"
	"ntl-test/backend/internal/handlers"
	"ntl-test/backend/internal/models"
	"ntl-test/backend/internal/repositories"
	"ntl-test/backend/internal/routes"
	"ntl-test/backend/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db := database.ConnectDatabase(
		database.Config{
			Host:     cfg.DBHost,
			Port:     cfg.DBPort,
			User:     cfg.DBUser,
			Password: cfg.DBPassword,
			Name:     cfg.DBName,
			SSLMode:  cfg.DBSSLMode,
		},
	)

	err := db.AutoMigrate(
		&models.User{},
		&models.Application{},
	)

	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	userRepository := repositories.NewUserRepository(db)
	applicationRepository := repositories.NewApplicationRepository(db)

	authService := services.NewAuthService(
		userRepository,
		cfg.JWTSecret,
	)

	applicationService := services.NewApplicationService(
		applicationRepository,
	)

	authHandler := handlers.NewAuthHandler(
		authService,
	)

	applicationHandler := handlers.NewApplicationHandler(
		applicationService,
	)

	router := gin.Default()

	routes.SetupRoutes(
		router,
		authHandler,
		applicationHandler,
		cfg.JWTSecret,
	)

	log.Printf("server running on :%s", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
