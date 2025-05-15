package controllers

import (
	"backend/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersService interface {
	Login(username string, password string) (int, string, error)
}

type UserController struct {
	usersService UsersService
}

func NewUserController(usersService UsersService) *UserController {
	return &UserController{
		usersService: usersService,
	}
}

func (c *UserController) Login(ctx *gin.Context) {
	var request domain.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, token, err := c.usersService.Login(request.Username, request.Password)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain.LoginResponse{
		UserID: userID,
		Token:  token,
	})
}
