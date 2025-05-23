package main

import (
	"backend/clients"
	"backend/controllers"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	var client = clients.NewMySQLClient()
	var userService = services.NewUsersService(client)
	userController := controllers.NewUserController(userService)

	// Endpoints para socios (punto 2) - No requieren autenticación según el enunciado
	router.GET("/activities", controllers.GetActivities)
	router.GET("/activities/:id", controllers.GetActivityByID)
	router.GET("/activities/search", controllers.SearchActivities)
	router.GET("/users/:userID/activities", controllers.GetUserActivities)
	router.POST("/users/:userID/enrollments", controllers.EnrollInActivity)

	router.POST("/login", userController.Login)

	router.Run(":8080")
}
