package main

import (
	"backend/clients"
	"backend/controllers"
	"backend/db"
	"backend/services"
	"backend/utils"

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

	// Inicializar la base de datos
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

	// Crear instancias de los middlewares
	authMiddleware := utils.AuthMiddleware(usersClient, "")       // Solo autenticación
	adminMiddleware := utils.AuthMiddleware(usersClient, "admin") // Requiere rol "admin"

	// Endpoints para socios (punto 2) - No requieren autenticación según el enunciado
	router.GET("/activities", actividadesController.GetActivities)
	router.GET("/activities/:id", actividadesController.GetActivityByID)
	router.GET("/activities/search", actividadesController.SearchActivities)
	// Requiere autenticación
	router.POST("/login", userController.Login)

	authgroup := router.Group("/")
	authgroup.Use(authMiddleware)
	{
		authgroup.GET("/users/:userID/activities", actividadesController.GetUserActivities)
		authgroup.POST("/users/:userID/enrollments", inscriptionsController.EnrollInActivity)
	}

	// Endpoints para administradores (requieren rol "admin")
	adminGroup := router.Group("/admin")
	adminGroup.Use(adminMiddleware)
	{
		adminGroup.POST("/activities", actividadesController.CreateActivity)
		adminGroup.PUT("/activities/:id", actividadesController.UpdateActivity)
		adminGroup.DELETE("/activities/:id", actividadesController.DeleteActivity)
	}

	router.Run(":8080")
}
