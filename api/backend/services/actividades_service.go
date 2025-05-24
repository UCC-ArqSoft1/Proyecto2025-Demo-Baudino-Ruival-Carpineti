package services

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

var mysqlClient = clients.NewMySQLClient()

type ActivitiesService struct {
	activitiesClient clients.ActivitiesClient
}

func NewActivitiesService(activitiesClient clients.ActivitiesClient) *ActivitiesService {
	return &ActivitiesService{
		activitiesClient: activitiesClient,
	}
}

// GetActivities returns all available activities
func (s *ActivitiesService) GetActivities() []domain.Activity {
	activitiesDAO, err := mysqlClient.GetAllActivities()
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
	activityDAO, err := mysqlClient.GetActivityByID(id)
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
	activitiesDAO, err := mysqlClient.GetAllActivities()
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
	return s.activitiesClient.GetActivitiesByUserID(userID)
}

// EnrollUserInActivity enrolls a user in a specific schedule
func (s *ActivitiesService) EnrollUserInActivity(userID, scheduleID int) error {
	// Get the schedule to check capacity
	var scheduleDAO dao.Schedules
	result := mysqlClient.DB.First(&scheduleDAO, scheduleID)
	if result.Error != nil {
		return fmt.Errorf("error getting schedule: %w", result.Error)
	}

	// Check if there's available capacity
	if scheduleDAO.Cupo <= 0 {
		return fmt.Errorf("no available capacity in this schedule")
	}

	// Check if user is already enrolled
	var existingEnrollment dao.Inscription
	result = mysqlClient.DB.Where("usuario_id = ? AND horario_id = ?", userID, scheduleID).First(&existingEnrollment)
	if result.Error == nil {
		return fmt.Errorf("user is already enrolled in this schedule")
	}

	// Create the enrollment
	enrollment := dao.Inscription{
		UsuarioID:        userID,
		HorarioID:        scheduleID,
		FechaInscripcion: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Start a transaction
	tx := mysqlClient.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("error starting transaction: %w", tx.Error)
	}

	// Create enrollment
	if err := tx.Create(&enrollment).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error creating enrollment: %w", err)
	}

	// Update schedule capacity
	if err := tx.Model(&dao.Schedules{}).Where("id = ?", scheduleID).
		Update("cupo", gorm.Expr("cupo - 1")).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating schedule capacity: %w", err)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

// GetActivities returns all available activities
func GetActivities() []domain.Activity {
	activitiesDAO, err := mysqlClient.GetAllActivities()
	if err != nil {
		fmt.Printf("Error getting activities: %v\n", err)
		return []domain.Activity{}
	}

	// Convert from DAO to Domain
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
func GetActivityByID(id int) (domain.Activity, error) {
	activityDAO, err := mysqlClient.GetActivityByID(id)
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
func SearchActivities(category, keyword string) []domain.Activity {
	activitiesDAO, err := mysqlClient.GetAllActivities()
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
func GetActivitiesByUserID(userID int) []domain.Activity {
	// Here we'll implement the logic to get the activities a user is enrolled in
	// Including the specific schedules they're enrolled in
	return []domain.Activity{}
}

// EnrollUserInActivity enrolls a user in a specific schedule
func EnrollUserInActivity(userID, scheduleID int) error {
	// Get the schedule to check capacity
	var scheduleDAO dao.Schedules
	result := mysqlClient.DB.First(&scheduleDAO, scheduleID)
	if result.Error != nil {
		return fmt.Errorf("error getting schedule: %w", result.Error)
	}

	// Check if there's available capacity
	if scheduleDAO.Cupo <= 0 {
		return fmt.Errorf("no available capacity in this schedule")
	}

	// Check if user is already enrolled
	var existingEnrollment dao.Inscription
	result = mysqlClient.DB.Where("user_id = ? AND schedule_id = ?", userID, scheduleID).First(&existingEnrollment)
	if result.Error == nil {
		return fmt.Errorf("user is already enrolled in this schedule")
	}

	// Create the enrollment
	enrollment := dao.Inscription{
		UsuarioID:        userID,
		HorarioID:        scheduleID,
		FechaInscripcion: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Start a transaction
	tx := mysqlClient.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("error starting transaction: %w", tx.Error)
	}

	// Create enrollment
	if err := tx.Create(&enrollment).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error creating enrollment: %w", err)
	}

	// Update schedule capacity
	if err := tx.Model(&dao.Schedules{}).Where("id = ?", scheduleID).
		Update("cupo", gorm.Expr("cupo - 1")).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating schedule capacity: %w", err)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}
