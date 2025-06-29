package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/gorm"
)

// 1. Implementacion de interfaz
type ActividadesClient struct {
	db *gorm.DB
}

// 2. Constructor de la implementacion del cliente
func NewActividadesClient(db *gorm.DB) *ActividadesClient {
	return &ActividadesClient{
		db: db,
	}
}

// GetAllActivities obtiene todas las actividades
func (c *ActividadesClient) GetAllActivities() ([]dao.Activities, error) {
	var activities []dao.Activities
	// SELECT * FROM activities
	result := c.db.Preload("Horarios").Find(&activities)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting activities: %w", result.Error)
	}
	return activities, nil
}

// GetActivityByID obtiene una actividad por su ID
func (c *ActividadesClient) GetActivityByID(id int) (dao.Activities, error) {
	var activity dao.Activities
	// SELECT * FROM activities WHERE id = ?
	result := c.db.Preload("Horarios").First(&activity, id)
	if result.Error != nil {
		return dao.Activities{}, fmt.Errorf("error getting activity by ID: %w", result.Error)
	}
	return activity, nil
}
