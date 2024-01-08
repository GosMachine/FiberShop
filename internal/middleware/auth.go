package middleware

import (
	"FiberShop/internal/handlers"
	"FiberShop/internal/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

func (a *App) IsAuthenticated(c *fiber.Ctx) error {
	var IsAuthenticated bool
	email, token := utils.IsUserLoggedIn(c.Cookies("token"), a.log)
	if email != "" {
		IsAuthenticated = true
	}
	if IsAuthenticated && token != "" {
		handlers.SetCookie("token", token, c, time.Now().Add(time.Hour*336))
	}
	if !IsAuthenticated {
		return c.Redirect("/login")
	}
	return c.Next()
}
