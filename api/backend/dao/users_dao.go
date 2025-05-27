package dao


type User struct {
	ID            int           `gorm:"primaryKey"`
	PasswordHash  string        `gorm:"not null"`
	Email         string        `gorm:"not null;unique"`
	Rol           string        `gorm:"not null"` // "socio" o "admin"
	Username      string        `gorm:"unique"`
	Inscripciones []Inscription `gorm:"foreignKey:UsuarioID"`
}
