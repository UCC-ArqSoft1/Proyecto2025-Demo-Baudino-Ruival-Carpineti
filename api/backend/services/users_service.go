package services

import (
	"backend/dao"
	"backend/domain"
	"backend/utils"
	"fmt"
)

// UsersService define la interfaz para el servicio de usuarios
type UsersService interface {
	Login(request domain.LoginRequest) (domain.LoginResponse, error)
}

// UsersServiceImpl implementa la interfaz UsersService
type UsersServiceImpl struct {
	usersClient UsersClient
}

// UsersClient define la interfaz para el cliente de usuarios
type UsersClient interface {
	GetUserByUsername(username string) (dao.User, error)
}

// NewUsersService crea una nueva instancia del servicio de usuarios
func NewUsersService(usersClient UsersClient) UsersService {
	return &UsersServiceImpl{
		usersClient: usersClient,
	}
}

// Login implementa la autenticación de usuarios
func (s *UsersServiceImpl) Login(request domain.LoginRequest) (domain.LoginResponse, error) {
	userDAO, err := s.usersClient.GetUserByUsername(request.Username)
	if err != nil {
		return domain.LoginResponse{}, fmt.Errorf("error getting user: %w", err)
	}

	if utils.HashSHA256(request.Password) != userDAO.PasswordHash {
		return domain.LoginResponse{}, fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateJWT(userDAO.ID)
	if err != nil {
		return domain.LoginResponse{}, fmt.Errorf("error generating token: %w", err)
	}

	return domain.LoginResponse{
		UserID: userDAO.ID,
		Token:  token,
	}, nil
}
