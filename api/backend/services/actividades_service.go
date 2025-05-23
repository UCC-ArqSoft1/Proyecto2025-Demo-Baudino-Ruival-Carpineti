package services

import (
	"backend/clients"
	"backend/domain"
	"fmt"
)

var mysqlClient = clients.NewMySQLClient()

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
			Capacity:    activityDAO.Cupo,
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
		Capacity:    activityDAO.Cupo,
		Image:       activityDAO.Imagen,
		Status:      activityDAO.Estado,
		Schedules:   schedules,
	}, nil
}

// SearchActivities searches activities by category or keyword
func SearchActivities(category, keyword string) []domain.Activity {
	// Here we'll implement the logic to search activities
	// By category or keyword in title/description
	return []domain.Activity{}
}

// GetActivitiesByUserID returns the activities a user is enrolled in
func GetActivitiesByUserID(userID int) []domain.Activity {
	// Here we'll implement the logic to get the activities a user is enrolled in
	// Including the specific schedules they're enrolled in
	return []domain.Activity{}
}

// EnrollUserInActivity enrolls a user in a specific schedule
func EnrollUserInActivity(userID, scheduleID int) error {
	// Here we'll implement the logic to:
	// 1. Validate that there's available capacity in the schedule
	// 2. Create the enrollment with the current date
	// 3. Update the schedule's capacity
	return nil
}
