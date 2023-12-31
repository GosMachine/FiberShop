package routes

import (
	"FiberShop/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupPublicRoutes(a *fiber.App) {
	a.Get("/", handlers.HandleHome)
	a.Get("/login", handlers.HandleLogin)
	a.Get("/register", handlers.HandleRegister)
}
