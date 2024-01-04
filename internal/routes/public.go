package routes

import (
	"FiberShop/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupPublicRoutes(a *fiber.App, handle *handlers.Handle) {
	a.Get("/", handle.HandleHome)
	//a.Post("/email", handle.HandleEmail)
	authRoutes(a, handle)
}

func authRoutes(a *fiber.App, handle *handlers.Handle) {
	a.Get("/login", handle.HandleLogin)
	a.Post("/login", handle.HandleLoginForm)
	a.Get("/register", handle.HandleRegister)
	a.Post("/register", handle.HandleRegisterForm)
}
