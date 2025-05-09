// Get Hotel by ID

package main

import (
	"backend/controllers"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Endpoints para socios (punto 2) - No requieren autenticación según el enunciado
	router.GET("/actividades", controllers.GetActividades)
	router.GET("/actividades/buscar", controllers.BuscarActividades)
	router.GET("/actividades/:id", controllers.GetActividadByID)
	router.GET("/usuarios/:usuarioID/actividades", controllers.GetActividadesUsuario)
	router.POST("/usuarios/:usuarioID/inscripciones", controllers.InscribirEnActividad)

	services.Login("emiliano", "1234")

	router.Run(":8080")
}
