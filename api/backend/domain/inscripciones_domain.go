package domain

type Inscripcion struct {
	ID               int    `json:"id" gorm:"primaryKey"`
	UsuarioID        int    `json:"usuario_id"`
	HorarioID        int    `json:"horario_id"`
	FechaInscripcion string `json:"fecha_inscripcion"`
}
