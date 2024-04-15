package middleware

import (
	"FiberShop/internal/pkg/jwt"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type App struct {
	log *zap.Logger
}

func New(logger *zap.Logger) *App {
	return &App{log: logger}
}

func (a *App) Logger(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	a.log.Info(fmt.Sprintf("[%s] %s - %s\n", c.Method(), c.Path(), time.Since(start)))
	return err
}

func (a *App) IsAuthenticated(c *fiber.Ctx) error {
	if jwt.IsTokenValid(c.Cookies("token")) != "" {
		return c.Redirect("/login")
	}
	return c.Next()
}
