package domain

// Activity representa una actividad deportiva disponible en el gimnasio
type Activity struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	Instructor  string     `json:"instructor"`
	Duration    int        `json:"duration"`
	Image       string     `json:"image,omitempty"`
	Status      string     `json:"status"` // "activo", "inactivo"
	Schedules   []Schedule `json:"schedules" gorm:"foreignKey:ActivityID"`
}

// CreateActivityRequest define la estructura para crear una nueva actividad
type CreateActivityRequest struct {
	Title       string          `json:"title" binding:"required"`
	Description string          `json:"description" binding:"required"`
	Category    string          `json:"category" binding:"required"`
	Instructor  string          `json:"instructor" binding:"required"`
	Duration    int             `json:"duration" binding:"required"`
	Image       string          `json:"image"`
	Status      string          `json:"status" binding:"required"`
	Schedules   []ScheduleInput `json:"schedules" binding:"required"`
}

// UpdateActivityRequest define la estructura para actualizar una actividad
type UpdateActivityRequest struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Category    string          `json:"category"`
	Instructor  string          `json:"instructor"`
	Duration    int             `json:"duration"`
	Image       string          `json:"image"`
	Status      string          `json:"status"`
	Schedules   []ScheduleInput `json:"schedules"`
}

// ScheduleInput define la estructura para los horarios en las solicitudes
type ScheduleInput struct {
	WeekDay   string `json:"week_day" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
	Capacity  int    `json:"capacity" binding:"required"`
}
