package handlers

import (
	"FiberShop/internal/utils"
	"FiberShop/web/view/account"
	"FiberShop/web/view/auth"
	"FiberShop/web/view/email"
	"FiberShop/web/view/layout"
	"context"
	"math/rand"
	"strconv"
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
	a.Log.Info("", zap.Bool("email", user.EmailVerified))
	return a.renderTemplate(c, account.Settings(user.EmailVerified, a.getData(c, "Settings")))
}

func (a *Handle) HandleAccountVerification(c *fiber.Ctx) error {
	email1 := c.FormValue("email")
	code := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	go func(email string) {
		a.Redis.Client.Set(a.Redis.Ctx, "verificationCode:"+email, code, time.Minute*10)
		a.sendEmail(email, code)
	}(email1)
	return a.renderTemplate(c, email.Show(email1, "email_verification", a.getData(c, "Email")))
}

func (a *Handle) HandleSettingsChangeEmail(c *fiber.Ctx) error {
	email1 := c.FormValue("email")
	newEmail := c.FormValue("newEmail")
	if _, err := a.Db.User(newEmail); err == nil {
		return c.SendString("Email already used.")
	}
	code := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	a.Redis.Client.Set(a.Redis.Ctx, "verificationCode:"+email1, code, time.Minute*10)
	a.Redis.Client.Set(a.Redis.Ctx, "change_email:"+email1, newEmail, time.Minute*10)
	go a.sendEmail(email1, code)
	c.Set("HX-Redirect", "/email?action=change_email&address="+email1)
	return c.SendStatus(200)
}

func (a *Handle) HandleSettingsChangePass(c *fiber.Ctx) error {
	email1 := c.FormValue("email")
	code := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	a.Redis.Client.Set(a.Redis.Ctx, "verificationCode:"+email, code, time.Minute*10)
	go a.sendEmail(email1, code)
	return c.Redirect("/email?action=change_pass&address=" + email1)
}

func (a *Handle) HandleChangePass(c *fiber.Ctx) error {
	email1, _ := utils.IsTokenValid(c.Cookies("token"))
	exist := a.Redis.Client.Exists(a.Redis.Ctx, "emailAccess:"+email1).Val()
	if exist != 1 {
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
	token, err := a.Client.ChangePass(context.Background(), data.Email, data.Password, c.IP())
	if err != nil {
		a.Log.Error("changePass api client error", zap.Error(err))
		return c.SendString("Internal error. Please try again.")
	}
	SetCookie("token", token, c, time.Now().Add(30*24*time.Hour))
	a.Redis.Client.Del(a.Redis.Ctx, "emailAccess:"+data.Email)
	a.Log.Info("password changed successfully", zap.String("email", data.Email))
	go a.sendEmail(data.Email, "password successfully changed")
	c.Set("HX-Redirect", "/account")
	return c.SendStatus(200)
}
