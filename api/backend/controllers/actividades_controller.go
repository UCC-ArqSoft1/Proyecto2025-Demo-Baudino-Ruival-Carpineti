package controllers

import (
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivitiesController struct {
	service *services.ActivitiesService
}

func NewActivitiesController(service *services.ActivitiesService) *ActivitiesController {
	return &ActivitiesController{service: service}
}

func (c *ActivitiesController) GetActivities(ctx *gin.Context) {
	activities := c.service.GetActivities()
	ctx.JSON(http.StatusOK, activities)
}

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

func (c *ActivitiesController) SearchActivities(ctx *gin.Context) {
	category := ctx.Query("category")
	keyword := ctx.Query("keyword")

	activities := c.service.SearchActivities(category, keyword)
	ctx.JSON(http.StatusOK, activities)
}

func (c *ActivitiesController) GetUserActivities(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	activities := c.service.GetActivitiesByUserID(userID)
	ctx.JSON(http.StatusOK, activities)
}

func (c *ActivitiesController) EnrollInActivity(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var request struct {
		ScheduleID int `json:"schedule_id" binding:"required"`
	}

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
