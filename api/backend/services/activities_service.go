package services

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"
	"fmt"
	"strings"
	"time"
)

// ActivitiesService define la interfaz para el servicio de actividades
type ActivitiesService interface {
	GetActivities() []domain.Activity
	GetActivityByID(id int) (domain.Activity, error)
	SearchActivities(category, keyword string) []domain.Activity
	GetActivitiesByUserID(userID int) []domain.Activity
	EnrollUserInActivity(userID, scheduleID int) error
}

// ActivitiesServiceImpl implementa la interfaz ActivitiesService
type ActivitiesServiceImpl struct {
	actividadesClient  *clients.ActividadesClient
	inscriptionsClient *clients.InscriptionsClient
}

// NewActivitiesService crea una nueva instancia del servicio de actividades
func NewActivitiesService(actividadesClient *clients.ActividadesClient, inscriptionsClient *clients.InscriptionsClient) ActivitiesService {
	return &ActivitiesServiceImpl{
		actividadesClient:  actividadesClient,
		inscriptionsClient: inscriptionsClient,
	}
}

// GetActivities returns all available activities
func (s *ActivitiesServiceImpl) GetActivities() []domain.Activity {
	activitiesDAO, err := s.actividadesClient.GetAllActivities()
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
func (s *ActivitiesServiceImpl) GetActivityByID(id int) (domain.Activity, error) {
	activityDAO, err := s.actividadesClient.GetActivityByID(id)
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
func (s *ActivitiesServiceImpl) SearchActivities(category, keyword string) []domain.Activity {
	activitiesDAO, err := s.actividadesClient.GetAllActivities()
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
func (s *ActivitiesServiceImpl) GetActivitiesByUserID(userID int) []domain.Activity {
	// Get user's inscriptions
	inscriptions, err := s.inscriptionsClient.GetUserInscriptions(userID)
	if err != nil {
		fmt.Printf("Error getting user inscriptions: %v\n", err)
		return []domain.Activity{}
	}

	// Get all activities
	activitiesDAO, err := s.actividadesClient.GetAllActivities()
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
func (s *ActivitiesServiceImpl) EnrollUserInActivity(userID, scheduleID int) error {
	// Verificar si el horario existe y tiene cupo disponible
	activity, err := s.actividadesClient.GetActivityByID(scheduleID)
	if err != nil {
		return fmt.Errorf("error getting activity: %v", err)
	}

	// Verificar si el horario tiene cupo disponible
	var schedule *domain.Schedule
	for _, s := range activity.Horarios {
		if s.ID == scheduleID {
			schedule = &domain.Schedule{
				ID:         s.ID,
				ActivityID: s.ActividadID,
				WeekDay:    s.DiaSemana,
				StartTime:  s.HoraInicio,
				EndTime:    s.HoraFin,
				Capacity:   s.Cupo,
			}
			break
		}
	}

	if schedule == nil {
		return fmt.Errorf("schedule not found")
	}

	// Verificar si el usuario ya está inscrito en este horario
	exists, err := s.inscriptionsClient.CheckExistingEnrollment(userID, scheduleID)
	if err != nil {
		return fmt.Errorf("error checking enrollment: %v", err)
	}
	if exists {
		return fmt.Errorf("user is already enrolled in this schedule")
	}

	// Crear la inscripción
	enrollment := dao.Inscription{
		UsuarioID:        userID,
		HorarioID:        scheduleID,
		FechaInscripcion: time.Now().Format("2006-01-02 15:04:05"),
	}

	/*
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

		func (s *ServicioActividades) FinalizarActividad(tx *gorm.DB, enrollment domain.Enrollment) error {
			// Intentar hacer commit
			if err := tx.Commit().Error; err != nil {
				return fmt.Errorf("error committing transaction: %w", err)
			}

			// Crear inscripción luego del commit exitoso
			err := s.inscriptionsClient.CreateEnrollment(enrollment)
	*/

	err = s.inscriptionsClient.CreateEnrollment(enrollment)
	if err != nil {
		return fmt.Errorf("error creating enrollment: %v", err)
	}

	return nil
}
