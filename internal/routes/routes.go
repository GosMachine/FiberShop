package routes

import (
	"FiberShop/internal/handlers"
	"FiberShop/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(a *fiber.App, handle *handlers.Handle, middle *middleware.App) {
	setupPublicRoutes(a, handle)
	setupProtectedRoutes(a, middle, handle)
}
