package domain

type Inscription struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	UserID         int    `json:"user_id"`
	ScheduleID     int    `json:"schedule_id"`
	EnrollmentDate string `json:"enrollment_date"`
}
