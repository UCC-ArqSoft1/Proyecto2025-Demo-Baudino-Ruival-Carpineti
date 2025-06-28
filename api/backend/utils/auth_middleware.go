package utils

import (
	"backend/clients"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifica el token JWT y el rol del usuario.
// Si requiredRole es "", solo verifica autenticación (sin validar rol).
func AuthMiddleware(usersClient *clients.UsersClient, requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "No autorizado"})
			return
		}

		// Validar el token JWT usando la función existente en utils
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token inválido o expirado"})
			return
		}

		// Extraer userID y rol del token
		userIDStr, ok := claims["jti"].(string)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "ID de usuario inválido"})
			return
		}
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "ID de usuario inválido"})
			return
		}

		// Obtener el usuario para verificar el rol
		user, err := usersClient.GetUserByID(userID)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Usuario no encontrado"})
			return
		}

		// Verificar rol si se especificó uno requerido
		if requiredRole != "" && user.Rol != requiredRole {
			c.AbortWithStatusJSON(403, gin.H{"error": "No tienes permisos suficientes"})
			return
		}

		// Guardar datos del usuario en el contexto para uso posterior
		c.Set("userID", userID)
		c.Set("userRole", user.Rol)
		c.Next()
	}
}
