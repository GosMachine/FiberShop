package handlers

import (
	"FiberShop/web/view/account"
	"FiberShop/web/view/auth"
	"FiberShop/web/view/layout"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (a *Handle) HandleAccount(c *fiber.Ctx) error {
	return a.renderTemplate(c, account.Index(a.getData(c, "My Account")))
}

func (a *Handle) HandleAccountSettings(c *fiber.Ctx) error {
	email := a.Redis.GetToken(c.Cookies("token"))
	verified, _ := a.Redis.GetEmailVerifiedCache(email)
	return a.renderTemplate(c, account.Settings(verified, a.getData(c, "Settings")))
}

func (a *Handle) HandleAccountVerification(c *fiber.Ctx) error {
	email := c.FormValue("email")
	a.sendVerificationCode(email)
	return c.Redirect("/email?action=email_verification&address=" + email)
}

func (a *Handle) HandleSettingsChangeEmail(c *fiber.Ctx) error {
	email := c.FormValue("email")
	newEmail := c.FormValue("newEmail")
	if _, err := a.Redis.GetEmailVerifiedCache(newEmail); err == nil {
		return c.SendString("Email already used.")
	}
	a.Redis.Client.Set(a.Redis.Ctx, "change_email:"+email, newEmail, time.Minute*10)
	a.sendVerificationCode(email)
	c.Set("HX-Redirect", "/email?action=change_email&address="+email)
	return c.SendStatus(200)
}

func (a *Handle) HandleSettingsChangePass(c *fiber.Ctx) error {
	email := c.FormValue("email")
	a.sendVerificationCode(email)
	return c.Redirect("/email?action=change_pass&address=" + email)
}

func (a *Handle) HandleChangePass(c *fiber.Ctx) error {
	email := a.Redis.GetToken(c.Cookies("token"))
	access := a.Redis.Client.Exists(a.Redis.Ctx, "emailAccess:"+email).Val()
	if access == 0 {
		return a.renderTemplate(c, layout.NotFound(a.getData(c, "Page not found")))
	}
	return a.renderTemplate(c, auth.ChangePass(a.getData(c, "Change Password")))
}

func (a *Handle) HandleChangePassForm(c *fiber.Ctx) error {
	var data RequestData
	if err := c.BodyParser(&data); err != nil {
		a.Log.Error("bodyParse error", zap.Error(err))
		return c.SendString("Internal error. Please try again.")
	}
	if data.Password != data.ConfirmPassword {
		return c.SendString("Password mismatch.")
	}
	token, err := a.Grpc.Auth.ChangePass(context.Background(), data.Email, data.Password, c.IP(), c.Cookies("token"))
	if err != nil {
		a.Log.Error("changePass api client error", zap.Error(err))
		return c.SendString("Internal error. Please try again.")
	}
	c.Cookie(&fiber.Cookie{Name: "token", Secure: true, Value: token, Expires: time.Now().Add(a.Redis.GetTokenTTL(token))})
	a.Redis.Client.Del(a.Redis.Ctx, "emailAccess:"+data.Email)
	a.Log.Info("password changed successfully", zap.String("email", data.Email))
	go a.sendEmail(data.Email, "password successfully changed")
	c.Set("HX-Redirect", "/account")
	return c.SendStatus(200)
}
