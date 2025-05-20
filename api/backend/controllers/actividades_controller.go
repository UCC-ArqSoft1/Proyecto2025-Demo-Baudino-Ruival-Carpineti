package controllers

import (
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetActivities returns all available activities
func GetActivities(c *gin.Context) {
	activities := services.GetActivities()
	c.JSON(http.StatusOK, activities)
}

// GetActivityByID returns a specific activity by its ID
func GetActivityByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	activity, err := services.GetActivityByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

// SearchActivities searches activities by category or keyword
func SearchActivities(c *gin.Context) {
	category := c.Query("category")
	keyword := c.Query("keyword")

	activities := services.SearchActivities(category, keyword)
	c.JSON(http.StatusOK, activities)
}

// GetUserActivities returns the activities a user is enrolled in
func GetUserActivities(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	activities := services.GetActivitiesByUserID(userID)
	c.JSON(http.StatusOK, activities)
}

// EnrollInActivity enrolls a user in a specific schedule
func EnrollInActivity(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var request struct {
		ScheduleID int `json:"schedule_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.EnrollUserInActivity(userID, request.ScheduleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment successful"})
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
