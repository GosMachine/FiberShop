package middleware

import (
	"FiberShop/internal/database/redis"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type App struct {
	log   *zap.Logger
	redis *redis.Redis
}

func New(logger *zap.Logger, redis *redis.Redis) *App {
	return &App{log: logger, redis: redis}
}

func (a *App) Logger(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	a.log.Info(fmt.Sprintf("[%s] %s - %s\n", c.Method(), c.Path(), time.Since(start)))
	return err
}

func (a *App) IsAuthenticated(c *fiber.Ctx) error {
	if a.redis.GetToken(c.Cookies("token")) == "" {
		return c.Redirect("/login")
	}
	return c.Next()
}
