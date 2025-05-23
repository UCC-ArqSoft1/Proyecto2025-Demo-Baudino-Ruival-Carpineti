package controllers

import (
	"backend/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivitiesService interface {
	GetActivities() []domain.Activity
	GetActivityByID(id int) (domain.Activity, error)
	SearchActivities(category, keyword string) []domain.Activity
	GetActivitiesByUserID(userID int) []domain.Activity
	EnrollUserInActivity(userID, scheduleID int) error
}

type ActivitiesController struct {
	activitiesService ActivitiesService
}

func NewActivitiesController(activitiesService ActivitiesService) *ActivitiesController {
	return &ActivitiesController{
		activitiesService: activitiesService,
	}
}

// GetActivities returns all available activities
func (c *ActivitiesController) GetActivities(ctx *gin.Context) {
	activities := c.activitiesService.GetActivities()
	ctx.JSON(http.StatusOK, activities)
}

// GetActivityByID returns a specific activity by its ID
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

// SearchActivities searches activities by category or keyword
func (c *ActivitiesController) SearchActivities(ctx *gin.Context) {
	category := ctx.Query("category")
	keyword := ctx.Query("keyword")

	activities := c.activitiesService.SearchActivities(category, keyword)
	ctx.JSON(http.StatusOK, activities)
}

// GetUserActivities returns the activities a user is enrolled in
func (c *ActivitiesController) GetUserActivities(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	activities := c.activitiesService.GetActivitiesByUserID(userID)
	ctx.JSON(http.StatusOK, activities)
}

// EnrollInActivity enrolls a user in a specific schedule
func (c *ActivitiesController) EnrollInActivity(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var request domain.EnrollRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.activitiesService.EnrollUserInActivity(userID, request.ScheduleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Enrollment successful"})
}

/*
// Validar LoginUsuario valida el login de un usuario
func LoginUsuario(c *gin.Context) {
	var request struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.Login(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
*/
