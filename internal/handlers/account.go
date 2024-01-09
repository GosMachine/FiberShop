package handlers

import (
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

func (a *Handle) HandleSettingsChangePass(c *fiber.Ctx) error {
	pass := c.FormValue("password")
	email := c.FormValue("email")
	user, err := a.Redis.GetUserCache(email)
	if err != nil {
		a.Log.Error("error getting user", zap.Error(err))
		return a.renderTemplate(c, "account/settings", fiber.Map{"Title": "Settings", "Error": "InternalError"})
	}
	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(pass)); err != nil {
		a.Log.Error("password wrong", zap.Error(err))
		return a.renderTemplate(c, "account/settings", fiber.Map{"Title": "Settings", "Error": "InvalidCredentials"})
	}
	return a.renderTemplate(c, "account/change_pass", fiber.Map{"Title": "Change pass", "Email": email, "Action": "change_pass"})
}

func (a *Handle) HandleChangePassForm(c *fiber.Ctx) error {
	pass := c.FormValue("password")
	email := c.FormValue("email")
	action := c.FormValue("action")
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		a.Log.Error("failed to generate password hash", zap.Error(err))
	}
	user, err := a.Redis.GetUserCache(email)
	if err != nil {
		a.Log.Error("failed to get user cache", zap.Error(err))
	}
	user.PassHash = passHash
	user.LastLoginIp = c.IP()
	user.LastLoginDate = time.Now()
	err = a.Db.UpdateUser(user)
	if err != nil {
		a.Log.Error("failed to update user", zap.Error(err))
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
