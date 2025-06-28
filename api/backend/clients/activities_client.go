package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/gorm"
)

// ActividadesClient implementa la interfaz services.ActivitiesClient
type ActividadesClient struct {
	db *gorm.DB
}

// NewActividadesClient crea una nueva instancia del cliente de actividades
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

// CreateActivity crea una nueva actividad
func (c *ActividadesClient) CreateActivity(activity dao.Activities) error {
	result := c.db.Create(&activity)
	if result.Error != nil {
		return fmt.Errorf("error creating activity: %w", result.Error)
	}
	return nil
}

// UpdateActivity actualiza una actividad existente
func (c *ActividadesClient) UpdateActivity(activity dao.Activities) error {
	// Eliminar los horarios existentes antes de guardar los nuevos
	if err := c.db.Where("actividad_id = ?", activity.ID).Delete(&dao.Schedules{}).Error; err != nil {
		return fmt.Errorf("error deleting old schedules: %w", err)
	}
	// Guardar la actividad con los nuevos horarios
	result := c.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&activity)
	if result.Error != nil {
		return fmt.Errorf("error updating activity: %w", result.Error)
	}
	return nil
}

// DeleteActivity elimina una actividad por su ID
func (c *ActividadesClient) DeleteActivity(id int) error {
	// Primero eliminamos los horarios asociados
	if err := c.db.Where("actividad_id = ?", id).Delete(&dao.Schedules{}).Error; err != nil {
		return fmt.Errorf("error deleting schedules: %w", err)
	}

	// Luego eliminamos la actividad
	result := c.db.Delete(&dao.Activities{}, id)
	if result.Error != nil {
		return fmt.Errorf("error deleting activity: %w", result.Error)
	}
	return nil
}
