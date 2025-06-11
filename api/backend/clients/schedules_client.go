package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/gorm"
)

type SchedulesClient struct {
	db *gorm.DB
}

func NewSchedulesClient(db *gorm.DB) *SchedulesClient {
	return &SchedulesClient{
		db: db,
	}
}

func (c *SchedulesClient) GetScheduleByID(scheduleID int) (dao.Schedules, error) {
	var schedule dao.Schedules
	// SELECT * FROM schedules WHERE id = ?
	result := c.db.First(&schedule, scheduleID)
	if result.Error != nil {
		return dao.Schedules{}, fmt.Errorf("error getting schedule: %w", result.Error)
	}
	return schedule, nil
}

func (c *SchedulesClient) UpdateScheduleCapacity(scheduleID int) error {
	// UPDATE schedules SET cupo = cupo - 1 WHERE id = ?
	result := c.db.Model(&dao.Schedules{}).Where("id = ?", scheduleID).
		Update("cupo", gorm.Expr("cupo - 1"))
	if result.Error != nil {
		return fmt.Errorf("error updating schedule capacity: %w", result.Error)
	}
	return nil
}
