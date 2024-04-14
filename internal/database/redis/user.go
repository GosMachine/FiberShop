package redis

import (
	"FiberShop/internal/models"
	"FiberShop/internal/utils"
	"encoding/json"
	"time"

	"go.uber.org/zap"
)

func (r *Redis) SetUserCache(user models.User) error {
	encode, err := utils.EncodeUserData(user)
	if err != nil {
		return err
	}
	r.Client.Set(r.Ctx, "UserData:"+user.Email, encode, time.Hour)
	return nil
}

func (r *Redis) GetUserCache(email string) (models.User, error) {
	userData, err := r.Client.Get(r.Ctx, "UserData:"+email).Result()
	if err == nil && userData != "" {
		if err != nil {
			r.Log.Error("error get user data from cache", zap.Error(err))
		} else {
			decode, err := utils.DecodeUserData(userData)
			if err != nil {
				r.Log.Error("error decode user data", zap.Error(err))
			}
			r.Log.Info("userData from cache", zap.String("email", email))
			return decode, nil
		}
	}
	return r.getUserFromDb(email)
}

func (r *Redis) getUserFromDb(email string) (models.User, error) {
	var user models.User
	if err := r.Db.UserPreload("Cart.Product", email, &user); err != nil {
		r.Log.Error("error preload user", zap.Error(err))
		return models.User{}, err
	}

	if err := r.SetUserCache(user); err != nil {
		r.Log.Error("error set userCache", zap.Error(err))
	}
	r.Log.Info("userData from postgres", zap.String("email", email))
	return user, nil
}

func encodeUserData(user models.User) (string, error) {
	data, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func decodeUserData(data string) (models.User, error) {
	var user models.User
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
