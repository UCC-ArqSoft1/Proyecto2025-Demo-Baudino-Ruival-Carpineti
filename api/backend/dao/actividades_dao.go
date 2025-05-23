package dao

// Actividad representa una actividad deportiva en el gimnasio.
// Puede tener múltiples horarios asociados.
type Activities struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Titulo      string `gorm:"not null" json:"titulo"`
	Descripcion string `gorm:"type:text" json:"descripcion"`
	Categoria   string `gorm:"not null" json:"categoria"` // Ej: Yoga, Spinning, etc.
	Instructor  string `gorm:"not null" json:"instructor"`
	Duracion    int    `gorm:"not null" json:"duracion"`       // Duración en minutos
	Cupo        int    `gorm:"not null" json:"cupo"`           // Cupo general (si se usa)
	Imagen      string `json:"imagen,omitempty"`               // URL o ruta de la imagen
	Estado      string `gorm:"default:'activo'" json:"estado"` // Estado actual: "activo", "inactivo"

	// Relación uno a muchos con Horario
	Horarios []Schedules `gorm:"foreignKey:ActividadID;constraint:OnDelete:CASCADE;" json:"horarios"`
}
