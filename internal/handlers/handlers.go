package handlers

import (
	"FiberShop/internal/database/postgres"
	"FiberShop/internal/database/redis"
	"FiberShop/internal/transport/grpc/auth"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handle struct {
	Log    *zap.Logger
	Client *auth.Client
	Db     *postgres.Storage
	Redis  *redis.Redis
}

func New(log *zap.Logger, client *auth.Client, db *postgres.Storage, redis *redis.Redis) *Handle {
	return &Handle{Log: log, Client: client, Db: db, Redis: redis}
}

func SetCookie(name, value string, c *fiber.Ctx, expires time.Time) {
	cookie := fiber.Cookie{
		Name:    name,
		Secure:  true,
		Value:   value,
		Expires: expires,
	}
	c.Cookie(&cookie)
}
