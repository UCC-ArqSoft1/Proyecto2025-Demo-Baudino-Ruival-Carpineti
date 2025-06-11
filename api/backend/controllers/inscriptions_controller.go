package controllers

import (
	"backend/domain"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// InscriptionsController maneja las peticiones HTTP relacionadas con inscripciones
type InscriptionsController struct {
	inscriptionsService services.InscriptionsService
}

// NewInscriptionsController crea una nueva instancia del controlador de inscripciones
func NewInscriptionsController(inscriptionsService services.InscriptionsService) *InscriptionsController {
	return &InscriptionsController{
		inscriptionsService: inscriptionsService,
	}
}

// EnrollInActivity maneja la petición para inscribir a un usuario en una actividad
func (c *InscriptionsController) EnrollInActivity(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var request domain.EnrollmentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.inscriptionsService.EnrollUserInActivity(userID, request.ScheduleID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully enrolled in activity"})
}
