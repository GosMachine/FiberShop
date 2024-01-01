package routes

import (
	"FiberShop/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupPublicRoutes(a *fiber.App, auth *handlers.Handle) {
	a.Get("/", handlers.HandleHome)
	authRoutes(a, auth)
}

func authRoutes(a *fiber.App, auth *handlers.Handle) {
	a.Get("/login", handlers.HandleLogin)
	a.Post("/login", auth.HandleLoginForm)
	a.Get("/register", handlers.HandleRegister)
	a.Post("/register", auth.HandleRegisterForm)
}
