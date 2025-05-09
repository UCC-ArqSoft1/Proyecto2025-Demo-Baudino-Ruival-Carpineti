package services

import (
	"backend/domain"
	"fmt"
)

// GetActividades retorna todas las actividades disponibles
func GetActividades() []domain.Actividad {
	// Aquí se implementará la lógica para obtener actividades de la base de datos
	// Incluyendo sus horarios
	// return []domain.Actividad{}

	// Datos de prueba temporales
	actividades := []domain.Actividad{
		{
			ID:          1,
			Titulo:      "Yoga Matutino",
			Descripcion: "Clase de yoga para principiantes",
			Categoria:   "Yoga",
			Instructor:  "María García",
			Duracion:    60,
			Cupo:        20,
			Imagen:      "yoga.jpg",
			Estado:      "activo",
			Horarios: []domain.Horario{
				{
					ID:          1,
					ActividadID: 1,
					DiaSemana:   "Lunes",
					HoraInicio:  "09:00",
					HoraFin:     "10:00",
					Cupo:        15,
				},
				{
					ID:          2,
					ActividadID: 1,
					DiaSemana:   "Miércoles",
					HoraInicio:  "09:00",
					HoraFin:     "10:00",
					Cupo:        15,
				},
			},
		},
		{
			ID:          2,
			Titulo:      "Spinning Intenso",
			Descripcion: "Clase de spinning de alta intensidad",
			Categoria:   "Spinning",
			Instructor:  "Juan Pérez",
			Duracion:    45,
			Cupo:        15,
			Imagen:      "spinning.jpg",
			Estado:      "activo",
			Horarios: []domain.Horario{
				{
					ID:          3,
					ActividadID: 2,
					DiaSemana:   "Martes",
					HoraInicio:  "18:00",
					HoraFin:     "18:45",
					Cupo:        10,
				},
			},
		},
	}
	return actividades
}

// GetActividadByID retorna una actividad por su ID
func GetActividadByID(id int) (domain.Actividad, error) {
	// Datos de prueba temporales
	actividades := []domain.Actividad{
		{
			ID:          1,
			Titulo:      "Yoga Matutino",
			Descripcion: "Clase de yoga para principiantes",
			Categoria:   "Yoga",
			Instructor:  "María García",
			Duracion:    60,
			Cupo:        20,
			Imagen:      "yoga.jpg",
			Estado:      "activo",
			Horarios: []domain.Horario{
				{
					ID:          1,
					ActividadID: 1,
					DiaSemana:   "Lunes",
					HoraInicio:  "09:00",
					HoraFin:     "10:00",
					Cupo:        15,
				},
				{
					ID:          2,
					ActividadID: 1,
					DiaSemana:   "Miércoles",
					HoraInicio:  "09:00",
					HoraFin:     "10:00",
					Cupo:        15,
				},
			},
		},
		{
			ID:          2,
			Titulo:      "Spinning Intenso",
			Descripcion: "Clase de spinning de alta intensidad",
			Categoria:   "Spinning",
			Instructor:  "Juan Pérez",
			Duracion:    45,
			Cupo:        15,
			Imagen:      "spinning.jpg",
			Estado:      "activo",
			Horarios: []domain.Horario{
				{
					ID:          3,
					ActividadID: 2,
					DiaSemana:   "Martes",
					HoraInicio:  "18:00",
					HoraFin:     "18:45",
					Cupo:        10,
				},
			},
		},
	}

	// Buscar la actividad por ID
	for _, actividad := range actividades {
		if actividad.ID == id {
			return actividad, nil
		}
	}

	return domain.Actividad{}, fmt.Errorf("actividad no encontrada")
}

// BuscarActividades busca actividades por categoría o palabra clave
func BuscarActividades(categoria, palabraClave string) []domain.Actividad {
	// Aquí se implementará la lógica para buscar actividades
	// Por categoría o palabra clave en título/descripción
	return []domain.Actividad{}
}

// GetActividadesByUsuarioID retorna las actividades a las que un usuario está inscrito
func GetActividadesByUsuarioID(usuarioID int) []domain.Actividad {
	// Aquí se implementará la lógica para obtener las actividades a las que un usuario está inscrito
	// Incluyendo los horarios específicos a los que está inscrito
	return []domain.Actividad{}
}

// InscribirUsuarioEnActividad inscribe a un usuario en un horario específico
func InscribirUsuarioEnActividad(usuarioID, horarioID int) error {
	// Aquí se implementará la lógica para:
	// 1. Validar que exista cupo disponible en el horario
	// 2. Crear la inscripción con la fecha actual
	// 3. Actualizar el cupo del horario
	return nil
}
