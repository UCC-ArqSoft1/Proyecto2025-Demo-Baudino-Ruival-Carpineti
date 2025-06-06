package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/gorm"
)

type SchedulesClient struct {
	db *DBClient
}

func NewSchedulesClient(db *DBClient) *SchedulesClient {
	return &SchedulesClient{
		db: db,
	}
}

func (c *SchedulesClient) GetScheduleByID(scheduleID int) (dao.Schedules, error) {
	var schedule dao.Schedules
	result := c.db.DB.First(&schedule, scheduleID)
	if result.Error != nil {
		return dao.Schedules{}, fmt.Errorf("error getting schedule: %w", result.Error)
	}
	return schedule, nil
}

func (c *SchedulesClient) UpdateScheduleCapacity(scheduleID int) error {
	result := c.db.DB.Model(&dao.Schedules{}).Where("id = ?", scheduleID).
		Update("cupo", gorm.Expr("cupo - 1"))
	if result.Error != nil {
		return fmt.Errorf("error updating schedule capacity: %w", result.Error)
	}
	return nil
}
