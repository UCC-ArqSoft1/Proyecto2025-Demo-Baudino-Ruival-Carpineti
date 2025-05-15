package services

import (
	"backend/dao"
	"backend/utils"
	"fmt"
)

type UsersClient interface {
	GetUserByUsername(username string) (dao.User, error)
}

type UsersService struct {
	usersClient UsersClient
}

func NewUsersService(usersClient UsersClient) *UsersService {
	return &UsersService{
		usersClient: usersClient,
	}
}

func (s *UsersService) Login(username, password string) (int, string, error) { //devuelve ID(usuario), token generado y error
	userDAO, err := s.usersClient.GetUserByUsername(username)
	if err != nil {
		return 0, "", fmt.Errorf("error getting user: %w", err)
	}

	if utils.HashSHA256(password) != userDAO.PasswordHash {
		return 0, "", fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateJWT(userDAO.ID)
	if err != nil {
		return 0, "", fmt.Errorf("error generating token: %w", err)
	}
	return userDAO.ID, token, nil
}
