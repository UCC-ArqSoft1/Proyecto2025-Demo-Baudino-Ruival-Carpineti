package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/gorm"
)

// InscriptionsClient implementa la interfaz services.InscriptionsClient
type InscriptionsClient struct {
	db *gorm.DB
}

// NewInscriptionsClient crea una nueva instancia del cliente de inscripciones
func NewInscriptionsClient(db *gorm.DB) *InscriptionsClient {
	return &InscriptionsClient{
		db: db,
	}
}

// GetUserInscriptions obtiene las inscripciones de un usuario
func (c *InscriptionsClient) GetUserInscriptions(userID int) ([]dao.Inscription, error) {
	var inscriptions []dao.Inscription
	// SELECT * FROM inscriptions WHERE usuario_id = ?
	result := c.db.Where("usuario_id = ?", userID).Find(&inscriptions)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting user inscriptions: %w", result.Error)
	}
	return inscriptions, nil
}

// CheckExistingEnrollment verifica si un usuario ya está inscrito en un horario
func (c *InscriptionsClient) CheckExistingEnrollment(userID, scheduleID int) (bool, error) {
	var enrollment dao.Inscription
	// SELECT * FROM inscriptions WHERE usuario_id = ? AND horario_id = ?
	result := c.db.Where("usuario_id = ? AND horario_id = ?", userID, scheduleID).First(&enrollment)
	if result.Error == nil {
		return true, nil
	}
	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	}
	return false, fmt.Errorf("error checking enrollment: %w", result.Error)
}

// CreateEnrollment crea una nueva inscripción
func (c *InscriptionsClient) CreateEnrollment(enrollment dao.Inscription) error {
	// INSERT INTO inscriptions (usuario_id, horario_id, fecha_inscripcion) VALUES (?, ?, ?)
	result := c.db.Create(&enrollment)
	if result.Error != nil {
		return fmt.Errorf("error creating enrollment: %w", result.Error)
	}
	return nil
}
