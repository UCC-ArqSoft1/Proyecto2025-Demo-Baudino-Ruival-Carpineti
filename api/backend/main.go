package main

import (
	"backend/clients"
	"backend/controllers"
	"backend/db"
	"backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//var mysqlClient = clients.NewMySQLClient()

	// Middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//
	db := db.InitDB()

	// Inicializar clients específicos
	usersClient := clients.NewUsersClient(db)
	actividadesClient := clients.NewActividadesClient(db)
	schedulesClient := clients.NewSchedulesClient(db)
	inscriptionsClient := clients.NewInscriptionsClient(db)

	// Configuración de servicios y controladores de usuarios
	userService := services.NewUsersService(usersClient)
	userController := controllers.NewUserController(userService)

	// Configuración de servicios de inscripciones (debe crearse primero)
	inscriptionsService := services.NewInscriptionsService(inscriptionsClient, schedulesClient)
	inscriptionsController := controllers.NewInscriptionsController(inscriptionsService)

	// Configuración de servicios y controladores de actividades
	// actividadesClient implementa la interfaz services.ActivitiesClient
	actividadesService := services.NewActivitiesService(actividadesClient, inscriptionsService)
	actividadesController := controllers.NewActivitiesController(actividadesService)

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
