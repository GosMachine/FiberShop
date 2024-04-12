package handlers

import (
	"FiberShop/internal/utils"
	"FiberShop/web/view/account"
	"FiberShop/web/view/auth"
	"FiberShop/web/view/email"
	"FiberShop/web/view/layout"
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
	go func(email string) {
		a.Redis.Client.Set(a.Redis.Ctx, "verificationCode:"+email, code, time.Minute*10)
		a.Redis.Client.Set(a.Redis.Ctx, "change_email:"+email, newEmail, time.Minute*10)
		a.sendEmail(email, code)
	}(email1)
	c.Set("HX-Redirect", "/email?action=change_email&address="+email1)
	return c.SendStatus(200)
}

func (a *Handle) HandleSettingsChangePass(c *fiber.Ctx) error {
	email1 := c.FormValue("email")
	code := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	go func(email string) {
		a.Redis.Client.Set(a.Redis.Ctx, "verificationCode:"+email, code, time.Minute*10)
		a.sendEmail(email, code)
	}(email1)
	return c.Redirect("/email?action=change_pass&address=" + email1)
}

func (a *Handle) HandleChangePass(c *fiber.Ctx) error {
	email1, _ := utils.IsTokenValid(c.Cookies("token"))
	exist := a.Redis.Client.Exists(a.Redis.Ctx, "emailVerified:"+email1).Val()
	if exist != 1 {
		return a.renderTemplate(c, layout.NotFound(a.getData(c, "Page not found")))
	}
	return a.renderTemplate(c, auth.ChangePass(a.getData(c, "Change Password")))
}

func (a *Handle) HandleChangePassForm(c *fiber.Ctx) error {
	pass := c.FormValue("password")
	confirmPassword := c.FormValue("confirmPassword")
	email1 := c.FormValue("email")
	if pass != confirmPassword || pass == "" || len(pass) < 8 {
		return c.SendString("Password does not match or less than 8 characters.")
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		a.Log.Error("failed to generate password hash", zap.Error(err))
		return c.SendString("Internal error. Please try again.")
	}
	user, err := a.Redis.GetUserCache(email1)
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
	a.Redis.Client.Del(a.Redis.Ctx, "emailVerified:"+email1)
	a.Log.Info("password changed successfully", zap.String("email", email1))
	c.Set("HX-Redirect", "/account")
	return c.SendStatus(200)
}
