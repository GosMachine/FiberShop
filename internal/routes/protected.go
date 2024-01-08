package routes

import (
	"FiberShop/internal/handlers"
	"FiberShop/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupProtectedRoutes(a *fiber.App, middle *middleware.App, handle *handlers.Handle) {
	accountRoutes(a, middle, handle)

}

func accountRoutes(a *fiber.App, middle *middleware.App, handle *handlers.Handle) {
	account := a.Group("/account")
	account.Use(middle.IsAuthenticated)
	account.Get("/logout", handlers.HandleLogout)
	account.Post("/change_pass", handle.HandleChangePassForm)
	account.Get("/", handle.HandleAccount)
}
