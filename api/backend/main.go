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

	// Middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	var mysqlClient = clients.NewMySQLClient()

	// Configuración de servicios y controladores de usuarios
	var userService = services.NewUsersService(mysqlClient)
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
