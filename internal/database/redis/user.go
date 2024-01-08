package redis

import (
	"FiberShop/internal/models"
	"FiberShop/internal/utils"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func (r *Redis) GetUserCache(email string) (models.User, error) {
	timeStart := time.Now()
	userData, err := r.Client.Get(r.Ctx, "UserData:"+email).Result()
	fmt.Println(time.Since(timeStart))
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
	return r.setUserCache(email)
}

func (r *Redis) setUserCache(email string) (models.User, error) {
	var (
		err  error
		user models.User
	)
	user, err = r.Db.User(email)
	if err != nil {
		r.Log.Error("error set user cache", zap.Error(err))
		return models.User{}, err
	}
	encode, err := utils.EncodeUserData(user)
	if err != nil {
		r.Log.Error("error encode user", zap.Error(err))
		return models.User{}, err
	}
	r.Client.Set(r.Ctx, "UserData:"+email, encode, time.Hour)
	r.Log.Info("userData from postgres", zap.String("email", email))
	return user, nil
}
