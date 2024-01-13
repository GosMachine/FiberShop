package redis

import (
	"go.uber.org/zap"
	"time"
)

func (r *Redis) isIPCounted(key string) bool {
	return r.Client.Get(r.Ctx, key).Err() == nil
}

func (r *Redis) IncrementViewCounter(url, key string) {
	if r.isIPCounted(key) {
		return
	}
	err := r.Client.Incr(r.Ctx, "viewers:"+url).Err()
	if err != nil {
		r.Log.Error("error increment view counter", zap.Error(err))
		return
	}
	r.setUniqueIp(key)
}

func (r *Redis) setUniqueIp(key string) {
	err := r.Client.Set(r.Ctx, key, 1, 24*time.Hour).Err()
	if err != nil {
		r.Log.Error("error set unique ip", zap.Error(err))
	}
}
