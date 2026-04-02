// Package helper маленькие методы которые могут применяться по всему проекту
package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

// GenerateJWT генерация JWT токена для пользователя
func GenerateJWT(userID string, email string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default-secret-key"
	}

	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"realm_access": map[string]interface{}{
			"roles": []string{"role_user"},
		},
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("ошибка при генерации токена: %w", err)
	}

	return tokenString, nil
}
