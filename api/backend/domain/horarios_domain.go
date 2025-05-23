package domain

type Schedule struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	ActivityID int    `json:"activity_id"`
	WeekDay    string `json:"week_day"`   // Lunes, Martes, etc.
	StartTime  string `json:"start_time"` // Formato: HH:MM
	EndTime    string `json:"end_time"`   // Formato: HH:MM
	Capacity   int    `json:"capacity"`   // Cupo específico para este horario
}
