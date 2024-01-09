package utils

import (
	gosjwt "github.com/GosMachine/ServiceAuth/pkg/jwt"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"os"
	"time"
)

func IsUserLoggedIn(token string, Log *zap.Logger) (string, string) {
	const op = "Auth.IsUserLoggedIn"

	log := Log.With(
		zap.String("op", op),
		zap.String("token", token),
	)

	log.Debug("checking IsUserLoggedIn")

	email, tokenData := IsTokenValid(token)
	if email != "" {
		var err error
		token, err = updateToken(tokenData)
		if err != nil {
			log.Error("fail to update token", zap.Error(err))
		}
	}
	log.Debug("checked IsUserLoggedIn", zap.String("is_user_logged_in", email))

	return email, token
}

func IsTokenValid(token string) (string, jwt.MapClaims) {
	tokenData, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return "", nil
	}
	claims := tokenData.Claims.(jwt.MapClaims)
	return claims["email"].(string), claims
}

func updateToken(claims jwt.MapClaims) (string, error) {
	email := claims["email"].(string)
	rememberMe := claims["remember"].(string)
	exp := time.Unix(int64(claims["exp"].(float64)), 0)
	if exp.Sub(time.Now()) <= 48*time.Hour && rememberMe == "on" {
		token, err := NewToken(email, rememberMe, time.Duration(time.Now().Add(time.Hour*336).Unix()))
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", nil
}

func NewToken(email, rememberMe string, duration time.Duration) (string, error) {
	return gosjwt.NewToken(email, rememberMe, duration)
}
