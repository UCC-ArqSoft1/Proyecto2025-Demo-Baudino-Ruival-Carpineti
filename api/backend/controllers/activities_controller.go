package controllers

import (
	"backend/domain"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ActivitiesController maneja las peticiones HTTP relacionadas con actividades
type ActivitiesController struct {
	service services.ActivitiesService
}

// NewActivitiesController crea una nueva instancia del controlador de actividades
func NewActivitiesController(service services.ActivitiesService) *ActivitiesController {
	return &ActivitiesController{service: service}
}

// GetActivities maneja la petición para obtener todas las actividades
func (c *ActivitiesController) GetActivities(ctx *gin.Context) {
	activities := c.service.GetActivities()
	ctx.JSON(http.StatusOK, activities)
}

// GetActivityByID maneja la petición para obtener una actividad por su ID
func (c *ActivitiesController) GetActivityByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	activity, err := c.service.GetActivityByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	ctx.JSON(http.StatusOK, activity)
}

// SearchActivities maneja la petición para buscar actividades
func (c *ActivitiesController) SearchActivities(ctx *gin.Context) {
	category := ctx.Query("category")
	keyword := ctx.Query("keyword")

	activities := c.service.SearchActivities(category, keyword)
	ctx.JSON(http.StatusOK, activities)
}

// GetUserActivities maneja la petición para obtener las actividades de un usuario
func (c *ActivitiesController) GetUserActivities(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	activities := c.service.GetActivitiesByUserID(userID)
	ctx.JSON(http.StatusOK, activities)
}

// EnrollInActivity maneja la petición para inscribir a un usuario en una actividad
func (c *ActivitiesController) EnrollInActivity(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var request domain.EnrollmentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.service.EnrollUserInActivity(userID, request.ScheduleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Enrollment successful"})
}
