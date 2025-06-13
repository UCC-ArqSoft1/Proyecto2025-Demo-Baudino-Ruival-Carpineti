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
	// 1. Obtener datos desde DAO
	userDAO, err := s.usersClient.GetUserByUsername(request.Username)
	if err != nil {
		return domain.LoginResponse{}, fmt.Errorf("error getting user: %w", err)
	}

	// 2. Convertir DAO a Domain
	userDomain := domain.User{
		ID:           userDAO.ID,
		Username:     userDAO.Username,
		PasswordHash: userDAO.PasswordHash,
	}

	// 3. Validar contraseña (método de domain.User)
	if !userDomain.ValidatePassword(request.Password) {
		return domain.LoginResponse{}, fmt.Errorf("invalid credentials")
	}

	// 4. Generar token
	token, err := utils.GenerateJWT(userDomain.ID)
	if err != nil {
		return domain.LoginResponse{}, fmt.Errorf("error generating token: %w", err)
	}

	// 5. Devolver respuesta (solo con datos necesarios)
	return domain.LoginResponse{
		UserID: userDomain.ID,
		Token:  token,
	}, nil
}
