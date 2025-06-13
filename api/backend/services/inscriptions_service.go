package services

import (
	"backend/dao"
	"backend/domain"
	"fmt"
	"time"
)

// InscriptionsService define la interfaz para el servicio de inscripciones
type InscriptionsService interface {
	EnrollUserInActivity(userID, scheduleID int) error
	GetUserInscriptions(userID int) ([]domain.Inscription, error)
}

// InscriptionsClient define la interfaz para el cliente de inscripciones
type InscriptionsClient interface {
	GetUserInscriptions(userID int) ([]dao.Inscription, error)
	CheckExistingEnrollment(userID, scheduleID int) (bool, error)
	CreateEnrollment(enrollment dao.Inscription) error
}

// SchedulesClient define la interfaz para el cliente de horarios
type SchedulesClient interface {
	GetScheduleByID(id int) (dao.Schedules, error)
	UpdateScheduleCapacity(id int) error
}

// InscriptionsServiceImpl implementa la interfaz InscriptionsService
type InscriptionsServiceImpl struct {
	inscriptionsClient InscriptionsClient
	schedulesClient    SchedulesClient
}

// NewInscriptionsService crea una nueva instancia del servicio de inscripciones
func NewInscriptionsService(inscriptionsClient InscriptionsClient, schedulesClient SchedulesClient) InscriptionsService {
	return &InscriptionsServiceImpl{
		inscriptionsClient: inscriptionsClient,
		schedulesClient:    schedulesClient,
	}
}

// EnrollUserInActivity inscribe a un usuario en una actividad
func (s *InscriptionsServiceImpl) EnrollUserInActivity(userID, scheduleID int) error {
	// Get the schedule to check capacity
	schedule, err := s.schedulesClient.GetScheduleByID(scheduleID)
	if err != nil {
		return fmt.Errorf("error getting schedule: %w", err)
	}

	// Check if there's available capacity
	if schedule.Cupo <= 0 {
		return fmt.Errorf("no available capacity in this schedule")
	}

	// Check if user is already enrolled
	exists, err := s.inscriptionsClient.CheckExistingEnrollment(userID, scheduleID)
	if err != nil {
		return fmt.Errorf("error checking enrollment: %w", err)
	}
	if exists {
		return fmt.Errorf("user is already enrolled in this schedule")
	}

	// Create the enrollment
	enrollment := dao.Inscription{
		UsuarioID:        userID,
		HorarioID:        scheduleID,
		FechaInscripcion: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Create enrollment
	if err := s.inscriptionsClient.CreateEnrollment(enrollment); err != nil {
		return fmt.Errorf("error creating enrollment: %w", err)
	}

	// Update schedule capacity
	if err := s.schedulesClient.UpdateScheduleCapacity(scheduleID); err != nil {
		return fmt.Errorf("error updating schedule capacity: %w", err)
	}

	return nil
}

// GetUserInscriptions obtiene las inscripciones de un usuario
func (s *InscriptionsServiceImpl) GetUserInscriptions(userID int) ([]domain.Inscription, error) {
	inscriptionsDAO, err := s.inscriptionsClient.GetUserInscriptions(userID)
	if err != nil {
		return nil, err
	}

	inscriptions := make([]domain.Inscription, len(inscriptionsDAO))
	for i, inscriptionDAO := range inscriptionsDAO {
		inscriptions[i] = domain.Inscription{
			ID:             inscriptionDAO.ID,
			UserID:         inscriptionDAO.UsuarioID,
			ScheduleID:     inscriptionDAO.HorarioID,
			EnrollmentDate: inscriptionDAO.FechaInscripcion,
		}
	}
	return inscriptions, nil
}
