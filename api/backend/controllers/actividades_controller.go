package controllers

import (
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetActividades retorna todas las actividades disponibles
func GetActividades(c *gin.Context) {
	actividades := services.GetActividades()
	c.JSON(http.StatusOK, actividades)
}

// GetActividadByID retorna una actividad específica por su ID
func GetActividadByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	actividad, err := services.GetActividadByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}

	c.JSON(http.StatusOK, actividad)
}

// BuscarActividades busca actividades por categoría o palabra clave
func BuscarActividades(c *gin.Context) {
	categoria := c.Query("categoria")
	palabraClave := c.Query("palabraClave")

	actividades := services.BuscarActividades(categoria, palabraClave)
	c.JSON(http.StatusOK, actividades)
}

// GetActividadesUsuario retorna las actividades a las que el usuario está inscrito
func GetActividadesUsuario(c *gin.Context) {
	usuarioID, err := strconv.Atoi(c.Param("usuarioID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	actividades := services.GetActividadesByUsuarioID(usuarioID)
	c.JSON(http.StatusOK, actividades)
}

// InscribirEnActividad inscribe a un usuario en un horario específico de una actividad
func InscribirEnActividad(c *gin.Context) {
	usuarioID, err := strconv.Atoi(c.Param("usuarioID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	var request struct {
		HorarioID int `json:"horario_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.InscribirUsuarioEnActividad(usuarioID, request.HorarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inscripción exitosa"})
}

/*
// Validar LoginUsuario valida el login de un usuario
func LoginUsuario(c *gin.Context) {
	var request struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.Login(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
*/
