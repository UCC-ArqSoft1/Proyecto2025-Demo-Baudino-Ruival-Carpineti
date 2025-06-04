package domain

// Actividad representa una actividad deportiva disponible en el gimnasio
type Activity struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	Instructor  string     `json:"instructor"`
	Duration    int        `json:"duration"` // podria ser calculado a partir de horainicio y horafin
	Image       string     `json:"image,omitempty"`
	Status      string     `json:"status"` // "activo", "inactivo"
	Schedules   []Schedule `json:"schedules" gorm:"foreignKey:ActivityID"`
	// Multiples horarios para la misma actividad
	// Cada horario tiene su propio cupo
	// Podemos buscar actividades por día y hora
}
