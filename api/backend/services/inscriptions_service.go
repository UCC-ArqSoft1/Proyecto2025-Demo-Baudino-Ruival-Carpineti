package services

import (
	"backend/dao"
	"backend/domain"
	"fmt"
	"time"
)

// 1. Definicion de interfaces
type InscriptionsService interface {
	EnrollUserInActivity(userID, scheduleID int) error
	GetUserInscriptions(userID int) ([]domain.Inscription, error)
}

type InscriptionsClient interface {
	GetUserInscriptions(userID int) ([]dao.Inscription, error)
	CheckExistingEnrollment(userID, scheduleID int) (bool, error)
	CreateEnrollment(enrollment dao.Inscription) error
}

type SchedulesClient interface {
	GetScheduleByID(id int) (dao.Schedules, error)
	UpdateScheduleCapacity(id int) error
}

// 2. Estructura que implementa la interfaz
type InscriptionsServiceImpl struct {
	inscriptionsClient InscriptionsClient
	schedulesClient    SchedulesClient
}

// 3. Constructor de la implementacion
func NewInscriptionsService(inscriptionsClient InscriptionsClient, schedulesClient SchedulesClient) InscriptionsService {
	return &InscriptionsServiceImpl{
		inscriptionsClient: inscriptionsClient,
		schedulesClient:    schedulesClient,
	}
}

// EnrollUserInActivity inscribe a un usuario en una actividad
func (s *InscriptionsServiceImpl) EnrollUserInActivity(userID, scheduleID int) error {
	// Toma el horario por ID
	schedule, err := s.schedulesClient.GetScheduleByID(scheduleID)
	if err != nil {
		return fmt.Errorf("error getting schedule: %w", err)
	}

	// Checkea cupos
	if schedule.Cupo <= 0 {
		return fmt.Errorf("no hay cupo disponible en este horario")
	}

	// Checkea si el usuario ya está inscrito en el horario
	exists, err := s.inscriptionsClient.CheckExistingEnrollment(userID, scheduleID)
	if err != nil {
		return fmt.Errorf("error checking enrollment: %w", err)
	}
	if exists {
		return fmt.Errorf("ya estas inscripto en este horario")
	}

	// Crea el objeto inscripcion
	enrollment := dao.Inscription{
		UsuarioID:        userID,
		HorarioID:        scheduleID,
		FechaInscripcion: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Crea la inscripción en la base de datos
	if err := s.inscriptionsClient.CreateEnrollment(enrollment); err != nil {
		return fmt.Errorf("error creating enrollment: %w", err)
	}

	// Actualiza el cupo del horario
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
