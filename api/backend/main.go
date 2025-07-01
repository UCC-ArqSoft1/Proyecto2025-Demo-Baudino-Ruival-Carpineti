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

	// Configuración de CORS para permitir requests desde el frontend React
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Inicialización de la base de datos
	db := db.InitDB()

	// Inicialización de los clients (acceso a datos)
	usersClient := clients.NewUsersClient(db)
	actividadesClient := clients.NewActividadesClient(db)
	schedulesClient := clients.NewSchedulesClient(db)
	inscriptionsClient := clients.NewInscriptionsClient(db)

	// Servicios y controladores de usuarios
	userService := services.NewUsersService(usersClient)
	userController := controllers.NewUserController(userService)

	// Servicios y controladores de inscripciones
	inscriptionsService := services.NewInscriptionsService(inscriptionsClient, schedulesClient)
	inscriptionsController := controllers.NewInscriptionsController(inscriptionsService)

	// Servicios y controladores de actividades deportivas
	actividadesService := services.NewActivitiesService(actividadesClient, inscriptionsService)
	actividadesController := controllers.NewActivitiesController(actividadesService)

	// Middlewares de autenticación y autorización
	authMiddleware := utils.AuthMiddleware(usersClient, "")       // Verifica JWT, usado para endpoints de socios
	adminMiddleware := utils.AuthMiddleware(usersClient, "admin") // Verifica JWT y rol admin, usado para endpoints de administradores

	// ===================== ENDPOINTS PARA SOCIOS =====================
	// Estos endpoints permiten:
	// - Buscar actividades disponibles
	// - Obtener detalle de actividad por ID
	// - Listar actividades en las que el usuario está inscripto
	// - Inscribirse en una actividad
	// No requieren validación de permisos por token (excepto inscripción y "mis actividades")
	router.GET("/activities", actividadesController.GetActivities)           // Listado y búsqueda de actividades (solo activas)
	router.GET("/activities/:id", actividadesController.GetActivityByID)     // Detalle de actividad
	router.GET("/activities/search", actividadesController.SearchActivities) // Búsqueda avanzada
	router.POST("/login", userController.Login)                              // Autenticación de usuarios (genera JWT)

	// Endpoints que requieren autenticación (socio logueado)
	authgroup := router.Group("/")
	authgroup.Use(authMiddleware)
	{
		authgroup.GET("/users/:userID/activities", actividadesController.GetUserActivities)   // "Mis actividades"
		authgroup.POST("/users/:userID/enrollments", inscriptionsController.EnrollInActivity) // Inscripción en actividad
	}

	// ===================== ENDPOINTS PARA ADMINISTRADORES =====================
	// Estos endpoints permiten:
	// - Crear, editar y eliminar actividades deportivas
	// - Listar todas las actividades (activas e inactivas)
	// - Los administradores también pueden acceder a las funcionalidades de socio
	// Requieren validación de token y rol admin
	adminGroup := router.Group("/admin")
	adminGroup.Use(adminMiddleware)
	{
		adminGroup.POST("/activities", actividadesController.CreateActivity)       // Crear actividad
		adminGroup.PUT("/activities/:id", actividadesController.UpdateActivity)    // Editar actividad
		adminGroup.DELETE("/activities/:id", actividadesController.DeleteActivity) // Eliminar actividad
		adminGroup.GET("/activities", actividadesController.GetAllActivitiesAdmin) // Listar todas las actividades (admin)
	}

	// ===================== SEGURIDAD =====================
	// - Todas las operaciones de escritura (crear, editar, eliminar) requieren JWT firmado
	// - Las contraseñas de los usuarios se almacenan con hash seguro (SHA256)

	router.Run(":8080")
}
