package domain

// Actividad representa una actividad deportiva disponible en el gimnasio
type Actividad struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Titulo      string    `json:"titulo"`
	Descripcion string    `json:"descripcion"`
	Categoria   string    `json:"categoria"`
	Instructor  string    `json:"instructor"`
	Duracion    int       `json:"duracion"` // podria ser calculado a partir de horainicio y horafin
	Cupo        int       `json:"cupo"`     //Si manejamos el cupo por Horarios[] no es necesario cupo en actividad
	Imagen      string    `json:"imagen,omitempty"`
	Estado      string    `json:"estado"` // "activo", "inactivo"
	Horarios    []Horario `json:"horarios" gorm:"foreignKey:ActividadID"`
	// Multiples horarios para la misma actividad
	// Cada horario tiene su propio cupo
	// Podemos buscar actividades por día y hora
}

type Horario struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	ActividadID int    `json:"actividad_id"`
	DiaSemana   string `json:"dia_semana"`  // Lunes, Martes, etc.
	HoraInicio  string `json:"hora_inicio"` // Formato: HH:MM
	HoraFin     string `json:"hora_fin"`    // Formato: HH:MM
	Cupo        int    `json:"cupo"`        // Cupo específico para este horario
}
