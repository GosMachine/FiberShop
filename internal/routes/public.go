package routes

import (
	"FiberShop/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupPublicRoutes(a *fiber.App, auth *handlers.Handle) {
	a.Get("/", auth.HandleHome)
	authRoutes(a, auth)
}

func authRoutes(a *fiber.App, auth *handlers.Handle) {
	a.Get("/login", auth.HandleLogin)
	a.Post("/login", auth.HandleLoginForm)
	a.Get("/register", auth.HandleRegister)
	a.Post("/register", auth.HandleRegisterForm)
}
