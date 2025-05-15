package dao

type User struct {
	ID            int           `gorm:"primaryKey"`
	Nombre        string        `gorm:"unique"`
	PasswordHash  string        `gorm:"not null"`
	Email         string        `gorm:"not null;unique"`
	Rol           string        `gorm:"not null"` // "socio" o "admin"
	Inscripciones []Inscription `gorm:"foreignKey:UsuarioID"`
}

type Inscription struct {
	ID               int    `gorm:"primaryKey"`
	UsuarioID        int    `gorm:"not null"`
	HorarioID        int    `gorm:"not null"`
	FechaInscripcion string `gorm:"not null"`
}
