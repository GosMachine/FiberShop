package middleware

import (
	"FiberShop/internal/database/redis"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	redis *redis.Redis
}

func New(redis *redis.Redis) *App {
	return &App{redis: redis}
}

func (a *App) IsAuthenticated(c *fiber.Ctx) error {
	if a.redis.GetToken(c.Cookies("token")) == "" {
		return c.Redirect("/login")
	}
	return c.Next()
}
