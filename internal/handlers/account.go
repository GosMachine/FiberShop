package handlers

import (
	"FiberShop/internal/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (a *Handle) HandleAccount(c *fiber.Ctx) error {
	return a.renderTemplate(c, "account/index", fiber.Map{"Title": "My Account"})
}

func (a *Handle) HandleAccountSettings(c *fiber.Ctx) error {
	return a.renderTemplate(c, "account/settings", fiber.Map{"Title": "Settings"})
}

func (a *Handle) HandleAccountCart(c *fiber.Ctx) error {
	email, _ := utils.IsTokenValid(c.Cookies("token"))
	user, err := a.Redis.GetUserCache(email)
	if err != nil {
		a.Log.Error("error get user cache", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalError"})
	}
	return a.renderTemplate(c, "account/cart", fiber.Map{"Title": "Your Cart", "CartItems": user.Cart.CartItems})
}

func (a *Handle) HandleAccountVerification(c *fiber.Ctx) error {
	email := c.FormValue("email")
	go func(email string) {
		a.sendEmail(email)
	}(email)
	return a.renderTemplate(c, "email", fiber.Map{"Title": "Email", "Email": email, "Action": "email_verification"})
}

func (a *Handle) HandleSettingsChangeEmail(c *fiber.Ctx) error {
	email := c.FormValue("email")
	newEmail := c.FormValue("newEmail")
	if _, err := a.Db.User(newEmail); err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "EmailAlreadyUsed"})
	}
	go func(email string) {
		a.Redis.Client.Set(a.Redis.Ctx, "change_email:"+email, newEmail, time.Minute*30)
		a.sendEmail(email)
	}(email)
	return a.renderTemplate(c, "email", fiber.Map{"Title": "Email", "Email": email, "Action": "change_email"})
}

func (a *Handle) HandleSettingsChangePass(c *fiber.Ctx) error {
	email := c.FormValue("email")
	go func(email string) {
		a.sendEmail(email)
	}(email)
	return a.renderTemplate(c, "email", fiber.Map{"Title": "Email", "Email": email, "Action": "change_pass"})
}

func (a *Handle) HandleChangePassForm(c *fiber.Ctx) error {
	pass := c.FormValue("password")
	email := c.FormValue("email")
	action := c.FormValue("action")
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		a.Log.Error("failed to generate password hash", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalError"})
	}
	user, err := a.Redis.GetUserCache(email)
	if err != nil {
		a.Log.Error("failed to get user cache", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalError"})
	}
	user.PassHash = passHash
	user.LastLoginIp = c.IP()
	user.LastLoginDate = time.Now()
	err = a.Db.UpdateUser(user)
	if err != nil {
		a.Log.Error("failed to update user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalError"})
	}
	if err := a.Redis.SetUserCache(user); err != nil {
		a.Log.Error("error set userCache", zap.Error(err))
	}
	a.Log.Info("password changed successfully", zap.String("email", email))
	if action == "change_pass" {
		return c.Redirect("/account")
	}
	return c.Redirect("/")
}
