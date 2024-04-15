package jwt

import (
	"os"

	"github.com/golang-jwt/jwt"
)

// returning email if token valid
func IsTokenValid(token string) string {
	tokenData, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil || !tokenData.Valid {
		return ""
	}
	claims, ok := tokenData.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}
	email, ok := claims["email"].(string)
	if !ok {
		return ""
	}
	return email
}
