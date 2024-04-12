package routes

import (
	"FiberShop/internal/handlers"
	"FiberShop/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupProtectedRoutes(a *fiber.App, middle *middleware.App, handle *handlers.Handle) {
	accountRoutes(a, middle, handle)
	couponRoutes(a, middle, handle)
}

func accountRoutes(a *fiber.App, middle *middleware.App, handle *handlers.Handle) {
	account := a.Group("/account")
	account.Use(middle.IsAuthenticated)
	account.Post("/logout", handlers.HandleLogout)
	account.Post("/change_pass", handle.HandleChangePassForm)
	account.Get("/", handle.HandleAccount)
	account.Get("/settings", handle.HandleAccountSettings)
	account.Post("/settings/change_pass", handle.HandleSettingsChangePass)
	account.Get("/settings/change_pass", handle.HandleChangePass)
	account.Post("/settings/change_email", handle.HandleSettingsChangeEmail)
	account.Post("/email_verification", handle.HandleAccountVerification)
	account.Get("/cart", handle.HandleAccountCart)
	account.Post("/cart/delete", handle.HandleDeleteItem)
}

func couponRoutes(a *fiber.App, middle *middleware.App, handle *handlers.Handle) {
	coupon := a.Group("/coupon")
	coupon.Post("/discount", handle.HandleDiscountCoupon)
	coupon.Post("/gift", handle.HandleGiftCoupon)
}
