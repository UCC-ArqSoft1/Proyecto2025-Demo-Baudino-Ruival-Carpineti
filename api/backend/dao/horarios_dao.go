package dao

// Horario representa un turno específico en el que se dicta una actividad.
// Cada horario tiene su propio cupo y horario definido.
type Schedules struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	ActividadID int    `gorm:"not null" json:"actividad_id"` // FK hacia Actividad
	DiaSemana   string `gorm:"not null" json:"dia_semana"`   // Ej: "Lunes"
	HoraInicio  string `gorm:"not null" json:"hora_inicio"`  // Formato: "HH:MM"
	HoraFin     string `gorm:"not null" json:"hora_fin"`     // Formato: "HH:MM"
	Cupo        int    `gorm:"not null" json:"cupo"`         // Cupo específico para este horario
}
