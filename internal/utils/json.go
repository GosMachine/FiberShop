package utils

import (
	"FiberShop/internal/models"
	"encoding/json"
)

func EncodeUserData(user models.User) (string, error) {
	data, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func DecodeUserData(data string) (models.User, error) {
	var user models.User
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
