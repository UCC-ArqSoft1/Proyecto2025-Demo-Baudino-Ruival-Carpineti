package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/gorm"
)

// SchedulesClient implementa la interfaz services.SchedulesClient
type SchedulesClient struct {
	db *gorm.DB
}

// NewSchedulesClient crea una nueva instancia del cliente de horarios
func NewSchedulesClient(db *gorm.DB) *SchedulesClient {
	return &SchedulesClient{
		db: db,
	}
}

// GetScheduleByID obtiene un horario por su ID
func (c *SchedulesClient) GetScheduleByID(id int) (dao.Schedules, error) {
	var schedule dao.Schedules
	// SELECT * FROM schedules WHERE id = ?
	result := c.db.First(&schedule, id)
	if result.Error != nil {
		return dao.Schedules{}, fmt.Errorf("error getting schedule: %w", result.Error)
	}
	return schedule, nil
}

// UpdateScheduleCapacity actualiza la capacidad de un horario
func (c *SchedulesClient) UpdateScheduleCapacity(id int) error {
	// UPDATE schedules SET cupo = cupo - 1 WHERE id = ?
	result := c.db.Model(&dao.Schedules{}).Where("id = ?", id).UpdateColumn("cupo", gorm.Expr("cupo - ?", 1))
	if result.Error != nil {
		return fmt.Errorf("error updating schedule capacity: %w", result.Error)
	}
	return nil
}
