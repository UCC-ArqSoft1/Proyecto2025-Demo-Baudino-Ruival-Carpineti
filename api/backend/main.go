package main

import (
	"backend/clients"
	"backend/controllers"
	"backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	var mysqlClient = clients.NewMySQLClient()

	// Middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Configuración de servicios y controladores de usuarios
	var userService = services.NewUsersService(mysqlClient)
	userController := controllers.NewUserController(userService)

	// Configuración de servicios y controladores de actividades
	activitiesClient := clients.NewMySQLActivitiesClient(mysqlClient)
	activitiesService := services.NewActivitiesService(activitiesClient)
	activitiesController := controllers.NewActivitiesController(activitiesService)

	// Endpoints para socios (punto 2) - No requieren autenticación según el enunciado
	router.GET("/activities", activitiesController.GetActivities)
	router.GET("/activities/:id", activitiesController.GetActivityByID)
	router.GET("/activities/search", activitiesController.SearchActivities)
	router.GET("/users/:userID/activities", activitiesController.GetUserActivities)
	router.POST("/users/:userID/enrollments", activitiesController.EnrollInActivity)

	// Endpoints para autenticacion de usuarios (punto 1) - Requiere autenticación
	router.POST("/login", userController.Login)

	router.Run(":8080")
}
