package routes

import (
	"FiberShop/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(a *fiber.App, auth *handlers.Handle) {
	setupPublicRoutes(a, auth)
}
