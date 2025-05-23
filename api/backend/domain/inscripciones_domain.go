package domain

type Inscription struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	UserID         int    `json:"user_id"`
	ScheduleID     int    `json:"schedule_id"`
	EnrollmentDate string `json:"enrollment_date"`
}

// EnrollRequest representa la estructura de la solicitud para inscribirse en una actividad
type EnrollRequest struct {
	ScheduleID int `json:"schedule_id" binding:"required"`
}
