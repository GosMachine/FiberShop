package handlers

import (
	"FiberShop/internal/utils"
	"FiberShop/web/view/alerts"
	"FiberShop/web/view/auth"
	"FiberShop/web/view/email"
	"FiberShop/web/view/layout"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (a *Handle) HandleEmailForm(c *fiber.Ctx) error {
	type emailForm struct {
		Code   string `json:"code"`
		Email  string `json:"email"`
		Action string `json:"action"`
	}
	var data emailForm
	if err := c.BodyParser(&data); err != nil {
		a.Log.Error("error parse email form", zap.Error(err))
		return c.SendString("Internal error. Please try again.")
	}
	code, err := a.Redis.Client.Get(a.Redis.Ctx, "verificationCode:"+data.Email).Result()
	if err != nil {
		return c.SendString("Confirmation code has expired. Please request a new one")
	}
	if data.Code != code {
		return c.SendString("Incorrect code.")
	}
	switch data.Action {
	case "email_verification":
		go func(email string) {
			a.emailVerification(email)
		}(data.Email)
	case "account_recovery":
		token, err := utils.NewToken(data.Email, "", time.Hour*24)
		if err != nil {
			a.Log.Error("error create newToken", zap.Error(err))
			return c.SendString("Internal error. Please try again.")
		}
		SetCookie("token", token, c, time.Now().Add(time.Hour*24))
		return a.renderTemplate(c, auth.ChangePass(data.Action, a.getData(c, "Change pass")))
	case "change_pass":
		return a.renderTemplate(c, auth.ChangePass(data.Action, a.getData(c, "Change pass")))
	case "change_email":
		newEmail, err := a.Redis.Client.Get(a.Redis.Ctx, "change_email:"+data.Email).Result()
		if err != nil {
			a.Log.Error("error get newEmail", zap.Error(err))
			return c.SendString("Internal error. Please try again.")
		}
		user, err := a.Redis.GetUserCache(data.Email)
		if err != nil {
			a.Log.Error("error getting user cache", zap.Error(err))
			return c.SendString("Internal error. Please try again.")
		}
		user.Email = newEmail
		err = a.Db.UpdateUser(user)
		if err != nil {
			a.Log.Error("error change email", zap.Error(err))
			return c.SendString("Internal error. Please try again.")
		}

		if err := a.Redis.SetUserCache(user); err != nil {
			a.Redis.Log.Error("error set userCache", zap.Error(err))
		}
		a.Redis.Client.Del(a.Redis.Ctx, "change_email:"+data.Email, "UserData:"+data.Email)
		token, err := utils.NewToken(newEmail, "on", time.Hour*336)
		if err != nil {
			a.Log.Error("error create newToken", zap.Error(err))
			return c.SendString("Internal error. Please try again.")
		}
		SetCookie("token", token, c, time.Now().Add(time.Hour*336))
	}
	a.Redis.Client.Del(a.Redis.Ctx, "verificationCode:"+data.Email)
	c.Set("HX-Redirect", "/account")
	return c.SendStatus(200)
}

func (a *Handle) emailVerification(email string) {
	user, err := a.Redis.GetUserCache(email)
	fmt.Println(email)
	if err != nil {
		a.Log.Error("error get user cache", zap.Error(err))
		return
	}
	user.EmailVerified = true
	err = a.Db.UpdateUser(user)
	if err != nil {
		a.Log.Error("error update user", zap.Error(err))
		return
	}
	if err := a.Redis.SetUserCache(user); err != nil {
		a.Log.Error("error set userCache", zap.Error(err))
	}
}

func (a *Handle) HandleEmailResend(c *fiber.Ctx) error {
	email := c.FormValue("email")
	code := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	go func(email string) {
		a.sendEmail(email, code)
	}(email)
	return a.renderTemplate(c, alerts.Success("Code", a.getData(c, "Email")))
}

func (a *Handle) HandleEmail(c *fiber.Ctx) error {
	email1 := c.Query("address")
	if a.Redis.Client.Get(a.Redis.Ctx, "verificationCode:"+email1).Val() != "" {
		action := c.Query("action")
		return a.renderTemplate(c, email.Show(action, a.getData(c, "Email")))
	}
	return a.renderTemplate(c, layout.NotFound(a.getData(c, "Page not found")))
}
