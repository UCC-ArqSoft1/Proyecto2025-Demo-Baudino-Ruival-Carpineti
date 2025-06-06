package clients

import (
	"backend/dao"
	"fmt"
)

type UsersClient struct {
	db *DBClient
}

func NewUsersClient(db *DBClient) *UsersClient {
	return &UsersClient{
		db: db,
	}
}

func (c *UsersClient) GetUserByUsername(username string) (dao.User, error) {
	var userDAO dao.User
	result := c.db.DB.First(&userDAO, "username = ?", username)
	if result.Error != nil {
		return dao.User{}, fmt.Errorf("error getting user: %w", result.Error)
	}
	return userDAO, nil
}
