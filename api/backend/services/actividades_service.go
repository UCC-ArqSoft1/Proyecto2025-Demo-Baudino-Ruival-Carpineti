package services

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"
	"fmt"
	"strings"
	"time"
)

type ActivitiesService struct {
	mysqlClient *clients.MySQLClient
}

func NewActivitiesService(mysqlClient *clients.MySQLClient) *ActivitiesService {
	return &ActivitiesService{
		mysqlClient: mysqlClient,
	}
}

// GetActivities returns all available activities
func (s *ActivitiesService) GetActivities() []domain.Activity {
	activitiesDAO, err := s.mysqlClient.GetAllActivities()
	if err != nil {
		fmt.Printf("Error getting activities: %v\n", err)
		return []domain.Activity{}
	}

	activities := make([]domain.Activity, len(activitiesDAO))
	for i, activityDAO := range activitiesDAO {
		schedules := make([]domain.Schedule, len(activityDAO.Horarios))
		for j, scheduleDAO := range activityDAO.Horarios {
			schedules[j] = domain.Schedule{
				ID:         scheduleDAO.ID,
				ActivityID: scheduleDAO.ActividadID,
				WeekDay:    scheduleDAO.DiaSemana,
				StartTime:  scheduleDAO.HoraInicio,
				EndTime:    scheduleDAO.HoraFin,
				Capacity:   scheduleDAO.Cupo,
			}
		}

		activities[i] = domain.Activity{
			ID:          activityDAO.ID,
			Title:       activityDAO.Titulo,
			Description: activityDAO.Descripcion,
			Category:    activityDAO.Categoria,
			Instructor:  activityDAO.Instructor,
			Duration:    activityDAO.Duracion,
			Image:       activityDAO.Imagen,
			Status:      activityDAO.Estado,
			Schedules:   schedules,
		}
	}

	return activities
}

// GetActivityByID returns an activity by its ID
func (s *ActivitiesService) GetActivityByID(id int) (domain.Activity, error) {
	activityDAO, err := s.mysqlClient.GetActivityByID(id)
	if err != nil {
		return domain.Activity{}, fmt.Errorf("error getting activity by ID: %v", err)
	}

	schedules := make([]domain.Schedule, len(activityDAO.Horarios))
	for i, scheduleDAO := range activityDAO.Horarios {
		schedules[i] = domain.Schedule{
			ID:         scheduleDAO.ID,
			ActivityID: scheduleDAO.ActividadID,
			WeekDay:    scheduleDAO.DiaSemana,
			StartTime:  scheduleDAO.HoraInicio,
			EndTime:    scheduleDAO.HoraFin,
			Capacity:   scheduleDAO.Cupo,
		}
	}

	return domain.Activity{
		ID:          activityDAO.ID,
		Title:       activityDAO.Titulo,
		Description: activityDAO.Descripcion,
		Category:    activityDAO.Categoria,
		Instructor:  activityDAO.Instructor,
		Duration:    activityDAO.Duracion,
		Image:       activityDAO.Imagen,
		Status:      activityDAO.Estado,
		Schedules:   schedules,
	}, nil
}

// SearchActivities searches activities by category or keyword
func (s *ActivitiesService) SearchActivities(category, keyword string) []domain.Activity {
	activitiesDAO, err := s.mysqlClient.GetAllActivities()
	if err != nil {
		fmt.Printf("Error getting activities: %v\n", err)
		return []domain.Activity{}
	}

	// Filter activities based on category and keyword
	var filteredActivities []dao.Activities
	for _, activity := range activitiesDAO {
		matchesCategory := category == "" || activity.Categoria == category
		matchesKeyword := keyword == "" ||
			strings.Contains(strings.ToLower(activity.Titulo), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(activity.Descripcion), strings.ToLower(keyword))

		if matchesCategory && matchesKeyword {
			filteredActivities = append(filteredActivities, activity)
		}
	}

	// Convert from DAO to Domain
	activities := make([]domain.Activity, len(filteredActivities))
	for i, activityDAO := range filteredActivities {
		schedules := make([]domain.Schedule, len(activityDAO.Horarios))
		for j, scheduleDAO := range activityDAO.Horarios {
			schedules[j] = domain.Schedule{
				ID:         scheduleDAO.ID,
				ActivityID: scheduleDAO.ActividadID,
				WeekDay:    scheduleDAO.DiaSemana,
				StartTime:  scheduleDAO.HoraInicio,
				EndTime:    scheduleDAO.HoraFin,
				Capacity:   scheduleDAO.Cupo,
			}
		}

		activities[i] = domain.Activity{
			ID:          activityDAO.ID,
			Title:       activityDAO.Titulo,
			Description: activityDAO.Descripcion,
			Category:    activityDAO.Categoria,
			Instructor:  activityDAO.Instructor,
			Duration:    activityDAO.Duracion,
			Image:       activityDAO.Imagen,
			Status:      activityDAO.Estado,
			Schedules:   schedules,
		}
	}

	return activities
}

// GetActivitiesByUserID returns the activities a user is enrolled in
func (s *ActivitiesService) GetActivitiesByUserID(userID int) []domain.Activity {
	// Get user's inscriptions
	inscriptions, err := s.mysqlClient.GetUserInscriptions(userID)
	if err != nil {
		fmt.Printf("Error getting user inscriptions: %v\n", err)
		return []domain.Activity{}
	}

	// Get all activities
	activitiesDAO, err := s.mysqlClient.GetAllActivities()
	if err != nil {
		fmt.Printf("Error getting activities: %v\n", err)
		return []domain.Activity{}
	}

	// Create a map of schedule IDs the user is enrolled in
	enrolledScheduleIDs := make(map[int]bool)
	for _, inscription := range inscriptions {
		enrolledScheduleIDs[inscription.HorarioID] = true
	}

	// Filter activities that have schedules the user is enrolled in
	var userActivities []domain.Activity
	for _, activityDAO := range activitiesDAO {
		// Check if any of the activity's schedules are in the user's inscriptions
		hasEnrolledSchedule := false
		var enrolledSchedules []domain.Schedule

		for _, scheduleDAO := range activityDAO.Horarios {
			if enrolledScheduleIDs[scheduleDAO.ID] {
				hasEnrolledSchedule = true
				enrolledSchedules = append(enrolledSchedules, domain.Schedule{
					ID:         scheduleDAO.ID,
					ActivityID: scheduleDAO.ActividadID,
					WeekDay:    scheduleDAO.DiaSemana,
					StartTime:  scheduleDAO.HoraInicio,
					EndTime:    scheduleDAO.HoraFin,
					Capacity:   scheduleDAO.Cupo,
				})
			}
		}

		if hasEnrolledSchedule {
			userActivities = append(userActivities, domain.Activity{
				ID:          activityDAO.ID,
				Title:       activityDAO.Titulo,
				Description: activityDAO.Descripcion,
				Category:    activityDAO.Categoria,
				Instructor:  activityDAO.Instructor,
				Duration:    activityDAO.Duracion,
				Image:       activityDAO.Imagen,
				Status:      activityDAO.Estado,
				Schedules:   enrolledSchedules,
			})
		}
	}

	return userActivities
}

// EnrollUserInActivity enrolls a user in a specific schedule
func (s *ActivitiesService) EnrollUserInActivity(userID, scheduleID int) error {
	// Get the schedule to check capacity
	schedule, err := s.mysqlClient.GetScheduleByID(scheduleID)
	if err != nil {
		return fmt.Errorf("error getting schedule: %w", err)
	}

	// Check if there's available capacity
	if schedule.Cupo <= 0 {
		return fmt.Errorf("no available capacity in this schedule")
	}

	// Check if user is already enrolled
	exists, err := s.mysqlClient.CheckExistingEnrollment(userID, scheduleID)
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

	// Start a transaction
	tx := s.mysqlClient.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("error starting transaction: %w", tx.Error)
	}

	// Create enrollment
	if err := s.mysqlClient.CreateEnrollment(enrollment); err != nil {
		tx.Rollback()
		return fmt.Errorf("error creating enrollment: %w", err)
	}

	// Update schedule capacity
	if err := s.mysqlClient.UpdateScheduleCapacity(scheduleID); err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating schedule capacity: %w", err)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}
