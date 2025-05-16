package domain

type Horario struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	ActividadID int    `json:"actividad_id"`
	DiaSemana   string `json:"dia_semana"`  // Lunes, Martes, etc.
	HoraInicio  string `json:"hora_inicio"` // Formato: HH:MM
	HoraFin     string `json:"hora_fin"`    // Formato: HH:MM
	Cupo        int    `json:"cupo"`        // Cupo específico para este horario
}
