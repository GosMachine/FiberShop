package routes

import (
	"FiberShop/internal/handlers"
	"FiberShop/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/shareed2k/goth_fiber"
)

type Routes struct {
	app    *fiber.App
	handle *handlers.Handle
	middle *middleware.App
}

func New(app *fiber.App, handle *handlers.Handle, middle *middleware.App) *Routes {
	return &Routes{app: app, handle: handle, middle: middle}
}

func (r *Routes) SetupRoutes() {
	r.setupPublicRoutes()
	r.setupProtectedRoutes()
}

func (r *Routes) setupPublicRoutes() {
	r.app.Get("/", r.handle.HandleHome)
	r.app.Get("/metrics", monitor.New())
	r.app.Post("/email/resend", r.handle.HandleEmailResend)
	r.app.Post("/email", r.handle.HandleEmailForm)
	r.app.Get("/email", r.handle.HandleEmail)
	r.app.Get("/account_recovery", r.handle.HandleAccountRecovery)
	r.app.Post("/account_recovery", r.handle.HandleAccountRecoveryForm)
	r.app.Get("/change_pass", r.handle.HandleChangePass)
	r.app.Post("/change_pass", r.handle.HandleChangePassForm)
	r.authRoutes()
}

func (r *Routes) authRoutes() {
	r.app.Get("/login", r.handle.HandleLogin)
	r.app.Post("/login", r.handle.HandleLoginForm)
	r.app.Get("/register", r.handle.HandleRegister)
	r.app.Post("/register", r.handle.HandleRegisterForm)
	r.app.Get("/contact", r.handle.HandleContact)
	r.app.Post("/contact", r.handle.HandleContactForm)
	r.app.Get("/login/:provider", goth_fiber.BeginAuthHandler)
	r.app.Get("/auth/callback/:provider", r.handle.HandleOAuthCallback)
}

func (r *Routes) setupProtectedRoutes() {
	r.accountRoutes()
	r.couponRoutes()
}

func (r *Routes) accountRoutes() {
	account := r.app.Group("/account")
	account.Use(r.middle.IsAuthenticated)
	account.Post("/logout", r.handle.HandleLogout)
	account.Post("/settings/change_pass", r.handle.HandleSettingsChangePass)
	account.Get("/", r.handle.HandleAccount)
	account.Get("/settings", r.handle.HandleAccountSettings)
	account.Post("/settings/change_email", r.handle.HandleSettingsChangeEmail)
	account.Post("/email_verification", r.handle.HandleAccountVerification)
	account.Get("/cart", r.handle.HandleAccountCart)
	account.Post("/cart/delete", r.handle.HandleDeleteItem)
}

func (r *Routes) couponRoutes() {
	coupon := r.app.Group("/coupon")
	coupon.Use(r.middle.IsAuthenticated)
	//todo middleware for coupons (valid check)
	coupon.Post("/discount", r.handle.HandleDiscountCoupon)
	coupon.Post("/gift", r.handle.HandleGiftCoupon)
}

//todo admin group
