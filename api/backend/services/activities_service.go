package services

import (
	"backend/dao"
	"backend/domain"
	"fmt"
	"strings"
)

// ActivitiesClient define la interfaz para el cliente de actividades
type ActivitiesClient interface {
	GetAllActivities() ([]dao.Activities, error)
	GetActivityByID(id int) (dao.Activities, error)
	CreateActivity(activity dao.Activities) error
	UpdateActivity(activity dao.Activities) error
	DeleteActivity(id int) error
}

// ActivitiesServiceImpl implementa la interfaz ActivitiesService
type ActivitiesServiceImpl struct {
	actividadesClient   ActivitiesClient
	inscriptionsService InscriptionsService // Inyección de dependencia
}

// NewActivitiesService crea una nueva instancia del servicio de actividades
func NewActivitiesService(actividadesClient ActivitiesClient, inscriptionsService InscriptionsService) *ActivitiesServiceImpl {
	return &ActivitiesServiceImpl{
		actividadesClient:   actividadesClient,
		inscriptionsService: inscriptionsService,
	}
}

// GetActivities returns all available activities
func (s *ActivitiesServiceImpl) GetActivities(userRole string) []domain.Activity {
	activitiesDAO, err := s.actividadesClient.GetAllActivities()
	if err != nil {
		fmt.Printf("Error getting activities: %v\n", err)
		return []domain.Activity{}
	}

	var filtered []dao.Activities
	if userRole == "admin" {
		filtered = activitiesDAO
	} else {
		for _, a := range activitiesDAO {
			if a.Estado == "activo" {
				filtered = append(filtered, a)
			}
		}
	}

	activities := make([]domain.Activity, len(filtered))
	for i, activityDAO := range filtered {
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
	// Usar el servicio de inscripciones para obtener las inscripciones del usuario
	inscriptions, err := s.inscriptionsService.GetUserInscriptions(userID)
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
		enrolledScheduleIDs[inscription.ScheduleID] = true
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

// Conversión de ScheduleInput (domain) a Schedules (dao)
func scheduleInputsToDao(schedules []domain.ScheduleInput, actividadID int) []dao.Schedules {
	schedulesDAO := make([]dao.Schedules, len(schedules))
	for i, s := range schedules {
		schedulesDAO[i] = dao.Schedules{
			ActividadID: actividadID, // Si ya tienes el ID, si no, dejar en 0 y GORM lo setea
			DiaSemana:   s.WeekDay,
			HoraInicio:  s.StartTime,
			HoraFin:     s.EndTime,
			Cupo:        s.Capacity,
		}
	}
	return schedulesDAO
}

// Crear una nueva actividad
func (s *ActivitiesServiceImpl) CreateActivity(req domain.CreateActivityRequest) error {
	activityDAO := dao.Activities{
		Titulo:      req.Title,
		Descripcion: req.Description,
		Categoria:   req.Category,
		Instructor:  req.Instructor,
		Duracion:    req.Duration,
		Imagen:      req.Image,
		Estado:      req.Status,
		Horarios:    scheduleInputsToDao(req.Schedules, 0), // 0 porque aún no existe el ID
	}
	return s.actividadesClient.CreateActivity(activityDAO)
}

// Actualizar una actividad existente
func (s *ActivitiesServiceImpl) UpdateActivity(id int, req domain.UpdateActivityRequest) error {
	activityDAO, err := s.actividadesClient.GetActivityByID(id)
	if err != nil {
		return err
	}
	if req.Title != "" {
		activityDAO.Titulo = req.Title
	}
	if req.Description != "" {
		activityDAO.Descripcion = req.Description
	}
	if req.Category != "" {
		activityDAO.Categoria = req.Category
	}
	if req.Instructor != "" {
		activityDAO.Instructor = req.Instructor
	}
	if req.Duration != 0 {
		activityDAO.Duracion = req.Duration
	}
	if req.Image != "" {
		activityDAO.Imagen = req.Image
	}
	if req.Status != "" {
		activityDAO.Estado = req.Status
	}
	if len(req.Schedules) > 0 {
		activityDAO.Horarios = scheduleInputsToDao(req.Schedules, id)
	}
	return s.actividadesClient.UpdateActivity(activityDAO)
}

// Eliminar una actividad
func (s *ActivitiesServiceImpl) DeleteActivity(id int) error {
	return s.actividadesClient.DeleteActivity(id)
}
