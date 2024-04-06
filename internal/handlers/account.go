package handlers

import (
	"FiberShop/internal/utils"
	"FiberShop/web/view/account"
	"FiberShop/web/view/email"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (a *Handle) HandleAccount(c *fiber.Ctx) error {
	return a.renderTemplate(c, account.Index(a.getData(c, "My Account")))
}

func (a *Handle) HandleAccountSettings(c *fiber.Ctx) error {
	email, _ := utils.IsTokenValid(c.Cookies("token"))
	user, err := a.Redis.GetUserCache(email)
	if err != nil {
		a.Log.Error("error getting user", zap.Error(err))
	}

	return a.renderTemplate(c, account.Settings(user.EmailVerified, a.getData(c, "Settings")))
}

func (a *Handle) HandleAccountCart(c *fiber.Ctx) error {
	email, _ := utils.IsTokenValid(c.Cookies("token"))
	user, err := a.Redis.GetUserCache(email)
	if err != nil {
		a.Log.Error("error get user cache", zap.Error(err))
	}
	var totalCartPrice float64
	for _, v := range user.Cart.CartItems {
		totalCartPrice += v.TotalPrice
	}

	return a.renderTemplate(c, account.Cart(a.getData(c, "Your Cart"), user.Cart.CartItems, totalCartPrice))
}

func (a *Handle) HandleAccountVerification(c *fiber.Ctx) error {
	email1 := c.FormValue("email")
	go func(email string) {
		a.sendEmail(email)
	}(email1)
	return a.renderTemplate(c, email.Show("email_verification", a.getData(c, "Email")))
}

func (a *Handle) HandleSettingsChangeEmail(c *fiber.Ctx) error {
	email := c.FormValue("email")
	newEmail := c.FormValue("newEmail")
	if _, err := a.Db.User(newEmail); err == nil {
		return c.SendString("Email already used.")
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
		return c.SendString("Internal error. Please try again.")
	}
	user, err := a.Redis.GetUserCache(email)
	if err != nil {
		a.Log.Error("failed to get user cache", zap.Error(err))
		return c.SendString("Internal error. Please try again.")
	}
	user.PassHash = passHash
	user.LastLoginIp = c.IP()
	user.LastLoginDate = time.Now()
	err = a.Db.UpdateUser(user)
	if err != nil {
		a.Log.Error("failed to update user", zap.Error(err))
		return c.SendString("Internal error. Please try again.")
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
