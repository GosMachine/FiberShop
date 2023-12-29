package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
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
