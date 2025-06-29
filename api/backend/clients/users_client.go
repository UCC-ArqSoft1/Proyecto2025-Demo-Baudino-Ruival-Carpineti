package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/gorm"
)

// 1. Implementacion de interfaz
type UsersClient struct {
	db *gorm.DB
}

// 2. Constructor de la implementacion del cliente
func NewUsersClient(db *gorm.DB) *UsersClient {
	return &UsersClient{
		db: db,
	}
}

func (c *UsersClient) GetUserByUsername(username string) (dao.User, error) {
	var userDAO dao.User
	result := c.db.First(&userDAO, "username = ?", username)
	if result.Error != nil {
		return dao.User{}, fmt.Errorf("error getting user: %w", result.Error)
	}
	return userDAO, nil
}
