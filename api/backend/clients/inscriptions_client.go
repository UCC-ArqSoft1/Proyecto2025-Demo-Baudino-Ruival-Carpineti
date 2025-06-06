package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/gorm"
)

type InscriptionsClient struct {
	db *gorm.DB
}

func NewInscriptionsClient(db *gorm.DB) *InscriptionsClient {
	return &InscriptionsClient{
		db: db,
	}
}

func (c *InscriptionsClient) GetUserInscriptions(userID int) ([]dao.Inscription, error) {
	var inscriptions []dao.Inscription
	result := c.db.Where("usuario_id = ?", userID).Find(&inscriptions)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting user inscriptions: %w", result.Error)
	}
	return inscriptions, nil
}

func (c *InscriptionsClient) CheckExistingEnrollment(userID, scheduleID int) (bool, error) {
	var enrollment dao.Inscription
	result := c.db.Where("usuario_id = ? AND horario_id = ?", userID, scheduleID).First(&enrollment)
	if result.Error == nil {
		return true, nil
	}
	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	}
	return false, fmt.Errorf("error checking enrollment: %w", result.Error)
}

func (c *InscriptionsClient) CreateEnrollment(enrollment dao.Inscription) error {
	result := c.db.Create(&enrollment)
	if result.Error != nil {
		return fmt.Errorf("error creating enrollment: %w", result.Error)
	}
	return nil
}
