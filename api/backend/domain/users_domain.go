package domain

import (
	"backend/utils"
)

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

// ValidatePassword compara un input con el hash almacenado.
// Recibe una función hasher para no acoplar el dominio a utils.
func (u *User) ValidatePassword(inputPassword string) bool {
	return u.PasswordHash == utils.HashSHA256(inputPassword) // Usa utils directamente
}
