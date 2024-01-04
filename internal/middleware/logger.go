package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func (a *App) Logger(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	a.log.Info(fmt.Sprintf("[%s] %s - %s\n", c.Method(), c.Path(), time.Since(start)))
	return err
}
