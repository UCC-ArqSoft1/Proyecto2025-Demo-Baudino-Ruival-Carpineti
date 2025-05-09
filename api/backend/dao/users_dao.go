package dao

type Usuario struct {
	ID             int           `gorm:"primaryKey"`
	Nombre         string        `gorm:"unique"`
	HashedPassword string        `gorm:"not null"`
	Email          string        `gorm:"not null;unique"`
	Rol            string        `gorm:"not null"` // "socio" o "admin"
	Inscripciones  []Inscripcion `gorm:"foreignKey:UsuarioID"`
}

type Inscripcion struct {
	ID               int    `gorm:"primaryKey"`
	UsuarioID        int    `gorm:"not null"`
	HorarioID        int    `gorm:"not null"`
	FechaInscripcion string `gorm:"not null"`
}
