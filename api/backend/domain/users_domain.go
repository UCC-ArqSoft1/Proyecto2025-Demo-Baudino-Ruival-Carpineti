package domain

// User represents a user in the gym system
type User struct {
	ID           int           `json:"id" gorm:"primaryKey"`
	Username     string        `json:"username"`
	Email        string        `json:"email" gorm:"unique"`
	PasswordHash string        `json:"-"`    // La contraseña no se muestra en el JSON
	Role         string        `json:"role"` // "socio" o "admin"
	Inscriptions []Inscription `json:"inscriptions" gorm:"foreignKey:UserID"`
	// Ver que usuarios estan inscriptos en qué horarios
	// Controlar el cupo por horario
	// Facilitar la búsqueda de actividades disponibles en un día y hora específicos
}
