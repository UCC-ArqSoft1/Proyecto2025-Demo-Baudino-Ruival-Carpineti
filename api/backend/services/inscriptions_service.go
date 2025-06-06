package services

import (
	"backend/clients"
	"backend/dao"
	"fmt"
	"time"
)

type InscriptionsService struct {
	inscriptionsClient *clients.InscriptionsClient
	schedulesClient    *clients.SchedulesClient
}

func NewInscriptionsService(inscriptionsClient *clients.InscriptionsClient, schedulesClient *clients.SchedulesClient) *InscriptionsService {
	return &InscriptionsService{
		inscriptionsClient: inscriptionsClient,
		schedulesClient:    schedulesClient,
	}
}

func (s *InscriptionsService) EnrollUserInActivity(userID, scheduleID int) error {
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

func (s *InscriptionsService) GetUserInscriptions(userID int) ([]dao.Inscription, error) {
	return s.inscriptionsClient.GetUserInscriptions(userID)
}
