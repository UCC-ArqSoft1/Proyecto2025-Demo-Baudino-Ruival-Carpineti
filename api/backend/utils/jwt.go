package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtDuration = time.Hour * 24
	jwtSecret   = "jwtSecret"
)

func GenerateJWT(userID int, role string) (string, error) {
	//Setear expiracion
	expirationTime := time.Now().Add(jwtDuration)

	// Construir los claims personalizados
	claims := jwt.MapClaims{
		"jti": fmt.Sprintf("%d", userID),
		"rol": role,
		"exp": expirationTime.Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": "backend",
		"sub": "auth",
	}

	//Crear el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Firmar el token
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("token inválido o expirado: %w", err)
	}
	return claims, nil
}
