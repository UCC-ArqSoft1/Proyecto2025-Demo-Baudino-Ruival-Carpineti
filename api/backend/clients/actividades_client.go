package clients

import (
	"backend/dao"
	"fmt"
)

type ActividadesClient struct {
	db *DBClient
}

func NewActividadesClient(db *DBClient) *ActividadesClient {
	return &ActividadesClient{
		db: db,
	}
}

func (c *ActividadesClient) GetAllActivities() ([]dao.Activities, error) {
	var activities []dao.Activities
	result := c.db.DB.Preload("Horarios").Find(&activities)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting activities: %w", result.Error)
	}
	return activities, nil
}

func (c *ActividadesClient) GetActivityByID(id int) (dao.Activities, error) {
	var activity dao.Activities
	result := c.db.DB.Preload("Horarios").First(&activity, id)
	if result.Error != nil {
		return dao.Activities{}, fmt.Errorf("error getting activity by ID: %w", result.Error)
	}
	return activity, nil
}
