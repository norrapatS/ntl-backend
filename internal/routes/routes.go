package routes

import (
	"ntl-test/backend/internal/handlers"

	"github.com/gin-gonic/gin"
	"ntl-test/backend/internal/middleware"
)

func SetupRoutes(
	router *gin.Engine,
	authHandler *handlers.AuthHandler,
	applicationHandler *handlers.ApplicationHandler,
	jwtSecret string,
	adminHandler *handlers.AdminHandler,
) {
	router.Use(middleware.Cors())

	api := router.Group("/api/v1")

	auth := api.Group("/auth")
	applications := api.Group("/applications")
	admin := api.Group("/admin")

	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	applications.Use(
		middleware.AuthMiddleware(jwtSecret),
	)
	applications.POST("", applicationHandler.Create)

	admin.Use(
		middleware.AuthMiddleware(jwtSecret),
		middleware.AdminMiddleware(),
	)
	admin.GET(
		"/applications",
		adminHandler.GetApplications,
	)

	admin.GET(
		"/applications/:id",
		adminHandler.GetApplicationByID,
	)
}
