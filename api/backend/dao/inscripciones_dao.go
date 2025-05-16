package dao

type Inscription struct {
	ID               int    `gorm:"primaryKey"`
	UsuarioID        int    `gorm:"not null"`
	HorarioID        int    `gorm:"not null"`
	FechaInscripcion string `gorm:"not null"`
}
