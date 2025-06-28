package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/gorm"
)

type UsersClient struct {
	db *gorm.DB
}

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

func (c *UsersClient) GetUserByID(id int) (dao.User, error) {
	var userDAO dao.User
	result := c.db.First(&userDAO, id)
	if result.Error != nil {
		return dao.User{}, fmt.Errorf("error getting user: %w", result.Error)
	}
	return userDAO, nil
}

func (c *UsersClient) CreateUser(user dao.User) error {
	result := c.db.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("error creating user: %w", result.Error)
	}
	return nil
}
