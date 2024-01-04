package routes

import (
	"FiberShop/internal/handlers"
	"FiberShop/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupProtectedRoutes(a *fiber.App, middle *middleware.App) {
	accountRoutes(a, middle)

}

func accountRoutes(a *fiber.App, middle *middleware.App) {
	account := a.Group("/account")
	account.Use(middle.IsAuthenticated)
	account.Get("/logout", handlers.HandleLogout)
}
