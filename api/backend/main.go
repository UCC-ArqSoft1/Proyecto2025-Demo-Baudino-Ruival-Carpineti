package main

import (
	"backend/clients"
	"backend/controllers"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	var mysqlClient = clients.NewMySQLClient()

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
