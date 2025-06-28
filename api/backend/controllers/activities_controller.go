package controllers

import (
	"backend/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ActivitiesService define la interfaz para el servicio de actividades
// Declarada aquí porque el controller es quien la usa
// Así se sigue la buena práctica recomendada
type ActivitiesService interface {
	GetActivities() []domain.Activity
	GetActivityByID(id int) (domain.Activity, error)
	SearchActivities(category, keyword string) []domain.Activity
	GetActivitiesByUserID(userID int) []domain.Activity
	CreateActivity(req domain.CreateActivityRequest) error
	UpdateActivity(id int, req domain.UpdateActivityRequest) error
	DeleteActivity(id int) error
}

// ActivitiesController maneja las peticiones HTTP relacionadas con actividades
type ActivitiesController struct {
	activitiesService ActivitiesService
}

// NewActivitiesController crea una nueva instancia del controlador de actividades
func NewActivitiesController(activitiesService ActivitiesService) *ActivitiesController {
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

func (c *ActivitiesController) CreateActivity(ctx *gin.Context) {
	var req domain.CreateActivityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.activitiesService.CreateActivity(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}

func (c *ActivitiesController) UpdateActivity(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var req domain.UpdateActivityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.activitiesService.UpdateActivity(id, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *ActivitiesController) DeleteActivity(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := c.activitiesService.DeleteActivity(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
