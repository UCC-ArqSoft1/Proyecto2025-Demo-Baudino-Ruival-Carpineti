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

<<<<<<< Updated upstream
	// Configuración de servicios y controladores de usuarios
	var userService = services.NewUsersService(mysqlClient)
	userController := controllers.NewUserController(userService)

	// Configuración de servicios y controladores de actividades
	activitiesClient := clients.NewMySQLActivitiesClient(mysqlClient)
	activitiesService := services.NewActivitiesService(activitiesClient)
	activitiesController := controllers.NewActivitiesController(activitiesService)
=======
	// Inicializar cliente de base de datos
	dbClient := clients.NewDBClient()

	// Inicializar clients específicos
	usersClient := clients.NewUsersClient(dbClient)
	actividadesClient := clients.NewActividadesClient(dbClient)
	schedulesClient := clients.NewSchedulesClient(dbClient)
	inscriptionsClient := clients.NewInscriptionsClient(dbClient)

	// Configuración de servicios y controladores de usuarios
	userService := services.NewUsersService(usersClient)
	userController := controllers.NewUserController(userService)

	// Configuración de servicios y controladores de actividades
	actividadesService := services.NewActivitiesService(actividadesClient, inscriptionsClient)
	actividadesController := controllers.NewActivitiesController(actividadesService)

	// Configuración de servicios y controladores de inscripciones
	inscriptionsService := services.NewInscriptionsService(inscriptionsClient, schedulesClient)
	inscriptionsController := controllers.NewInscriptionsController(inscriptionsService)
>>>>>>> Stashed changes

	// Endpoints para socios (punto 2) - No requieren autenticación según el enunciado
	router.GET("/activities", actividadesController.GetActivities)
	router.GET("/activities/:id", actividadesController.GetActivityByID)
	router.GET("/activities/search", actividadesController.SearchActivities)
	router.GET("/users/:userID/activities", actividadesController.GetUserActivities)
	router.POST("/users/:userID/enrollments", inscriptionsController.EnrollInActivity)

	// Endpoints para autenticacion de usuarios (punto 1) - Requiere autenticación
	router.POST("/login", userController.Login)

	router.Run(":8080")
}
