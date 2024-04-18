package redis

import (
	"context"
	"time"

	"go.uber.org/zap"
)

func (r *Redis) SetEmailVerifiedCache(email string, value bool) error {
	r.Client.Set(r.Ctx, "emailVerified:"+email, value, time.Hour*24)
	return nil
}

func (r *Redis) GetEmailVerifiedCache(email string) (bool, error) {
	verified, err := r.Client.Get(r.Ctx, "emailVerified:"+email).Bool()
	if err != nil {
		r.Log.Error("error get user data from cache", zap.Error(err))
		verified, err = r.AuthClient.EmailVerified(context.Background(), email)
		r.SetEmailVerifiedCache(email, verified)
		return verified, err
	}
	r.Log.Info("emailVerified from cache", zap.String("email", email))
	return verified, nil
}
