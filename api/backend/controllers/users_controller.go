package controllers

import (
	"backend/domain"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController maneja las peticiones HTTP relacionadas con usuarios
type UserController struct {
	usersService services.UsersService
}

// NewUserController crea una nueva instancia del controlador de usuarios
func NewUserController(usersService services.UsersService) *UserController {
	return &UserController{
		usersService: usersService,
	}
}

// Login maneja la petición de inicio de sesión
func (c *UserController) Login(ctx *gin.Context) {
	var request domain.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.usersService.Login(request)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
