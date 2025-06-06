package controllers

import (
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InscriptionsController struct {
	inscriptionsService *services.InscriptionsService
}

func NewInscriptionsController(inscriptionsService *services.InscriptionsService) *InscriptionsController {
	return &InscriptionsController{
		inscriptionsService: inscriptionsService,
	}
}

func (c *InscriptionsController) EnrollInActivity(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var request struct {
		ScheduleID int `json:"schedule_id" binding:"required"`
	}

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
