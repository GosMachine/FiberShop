package routes

import (
	"FiberShop/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
)

func setupPublicRoutes(a *fiber.App, handle *handlers.Handle) {
	a.Get("/", handle.HandleHome)
	a.Post("/email/resend", handle.HandleEmailResend)
	a.Post("/email", handle.HandleEmail)
	a.Get("/account_recovery", handle.HandleAccountRecovery)
	a.Post("/account_recovery", handle.HandleAccountRecoveryForm)
	authRoutes(a, handle)
}

func authRoutes(a *fiber.App, handle *handlers.Handle) {
	a.Get("/login", handle.HandleLogin)
	a.Post("/login", handle.HandleLoginForm)
	a.Get("/register", handle.HandleRegister)
	a.Post("/register", handle.HandleRegisterForm)
	a.Get("/contact", handle.HandleContact)
	a.Post("/contact", handle.HandleContactForm)
	a.Get("/login/:provider", goth_fiber.BeginAuthHandler)
	a.Get("/auth/callback/:provider", handle.HandleOAuthCallback)
}
