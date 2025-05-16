package domain

// Usuario representa un usuario del sistema del gimnasio
type Usuario struct {
	ID            int           `json:"id" gorm:"primaryKey"`
	Nombre        string        `json:"nombre"`
	Email         string        `json:"email" gorm:"unique"`
	Password      string        `json:"-"`   // La contraseña no se muestra en el JSON
	Rol           string        `json:"rol"` // "socio" o "admin"
	Inscripciones []Inscripcion `json:"inscripciones" gorm:"foreignKey:UsuarioID"`
	// Ver que usuarios estan inscriptos en qué horarios
	// Controlar el cupo por horario
	// Facilitar la búsqueda de actividades disponibles en un día y hora específicos
}
