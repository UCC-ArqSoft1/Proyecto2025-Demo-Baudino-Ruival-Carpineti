package controllers

import (
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ActivitiesController maneja las peticiones HTTP relacionadas con actividades
type ActivitiesController struct {
	activitiesService services.ActivitiesService
}

// NewActivitiesController crea una nueva instancia del controlador de actividades
func NewActivitiesController(activitiesService services.ActivitiesService) *ActivitiesController {
	return &ActivitiesController{
		activitiesService: activitiesService,
	}
}

// GetActivities maneja la petición para obtener todas las actividades
func (c *ActivitiesController) GetActivities(ctx *gin.Context) {
	activities := c.activitiesService.GetActivities()
	ctx.JSON(http.StatusOK, activities)
}

// GetActivityByID maneja la petición para obtener una actividad por su ID
func (c *ActivitiesController) GetActivityByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	activity, err := c.activitiesService.GetActivityByID(id)
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

	activities := c.activitiesService.SearchActivities(category, keyword)
	ctx.JSON(http.StatusOK, activities)
}

// GetUserActivities maneja la petición para obtener las actividades de un usuario
func (c *ActivitiesController) GetUserActivities(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	activities := c.activitiesService.GetActivitiesByUserID(userID)
	ctx.JSON(http.StatusOK, activities)
}
