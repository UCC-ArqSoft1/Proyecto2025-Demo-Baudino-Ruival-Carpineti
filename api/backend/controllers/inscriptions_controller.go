package controllers

import (
	"backend/domain"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 1. Uso de la interfaz, NO de la implementación
// Sabe que existe un servicio de inscripciones, pero no le importa como se implementa
type InscriptionsController struct {
	inscriptionsService services.InscriptionsService
}

// 2. Constructor de la implementación del controlador
func NewInscriptionsController(inscriptionsService services.InscriptionsService) *InscriptionsController {
	return &InscriptionsController{
		inscriptionsService: inscriptionsService,
	}
}

// 3. Llega una peticion HTTP
// Llama al metodo, independientemente de la implementación concreta
func (c *InscriptionsController) EnrollInActivity(ctx *gin.Context) {
	// Obtiene ID
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	// Obtiene el cuerpo de la solicitud
	var request domain.EnrollmentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verifica que el scheduleID sea válido
	if err := c.inscriptionsService.EnrollUserInActivity(userID, request.ScheduleID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Inscripcion exitosa"})
}
