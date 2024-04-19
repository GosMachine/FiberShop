package handlers

import (
	"FiberShop/web/view/alerts"
	"FiberShop/web/view/email"
	"FiberShop/web/view/layout"
	"context"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
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
		go a.emailVerification(data.Email)
	case "change_pass":
		a.Redis.Client.Set(a.Redis.Ctx, "emailAccess:"+data.Email, true, time.Minute*10)
		c.Set("HX-Redirect", "/change_pass")
		return c.SendStatus(200)
	case "change_email":
		//todo продолжить тут
		newEmail, err := a.Redis.Client.Get(context.Background(), "change_email:"+data.Email).Result()
		if err != nil {
			a.Log.Error("error get newEmail", zap.Error(err))
			return c.SendString("Internal error. Please try again.")
		}
		token, err := a.Client.ChangeEmail(context.Background(), data.Email, newEmail, c.Cookies("token"))
		if err != nil {
			a.Log.Error("error change email", zap.Error(err))
			return c.SendString("Internal error. Please try again.")
		}
		a.Redis.Client.Del(a.Redis.Ctx, "change_email:"+data.Email, "emailVerified:"+data.Email)
		c.Cookie(&fiber.Cookie{Name: "token", Secure: true, Value: token, Expires: time.Now().Add(a.Redis.GetTokenTTL(token))})
		go a.sendEmail(newEmail, "email successfully changed")
	}
	a.Redis.Client.Del(a.Redis.Ctx, "verificationCode:"+data.Email)
	c.Set("HX-Redirect", "/account")
	return c.SendStatus(200)
}

func (a *Handle) emailVerification(email string) {
	if err := a.Client.EmailVerify(context.Background(), email); err != nil {
		a.Log.Error("error verify email", zap.Error(err))
		return
	}
	if err := a.Redis.SetEmailVerifiedCache(email, true); err != nil {
		a.Log.Error("error set userCache", zap.Error(err))
	}
}

func (a *Handle) HandleEmailResend(c *fiber.Ctx) error {
	a.sendVerificationCode(c.FormValue("email"))
	alert := alerts.Alert{
		Name:    "Code",
		Message: "Successfully sent",
		Button:  "resendBtn",
	}
	return a.renderTemplate(c, alerts.Success(alert, a.getData(c, "Email")))
}

func (a *Handle) HandleEmail(c *fiber.Ctx) error {
	email1 := c.Query("address")
	if a.Redis.Client.Get(a.Redis.Ctx, "verificationCode:"+email1).Val() != "" {
		action := c.Query("action")
		return a.renderTemplate(c, email.Show(email1, action, a.getData(c, "Email")))
	}
	return a.renderTemplate(c, layout.NotFound(a.getData(c, "Page not found")))
}

func (a *Handle) sendEmail(email, msg string) {
	message := gomail.NewMessage()
	message.SetHeader("From", "support@fiber.shop")
	message.SetHeader("To", email)
	message.SetHeader("Subject", "FiberShop")
	message.SetBody("text/plain", msg)
	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_NAME"), os.Getenv("EMAIL_PASS"))
	if err := dialer.DialAndSend(message); err != nil {
		a.Log.Error("error send email "+email, zap.Error(err))
		return
	}
	a.Log.Info("send email success", zap.String("email", email))
}

func (a *Handle) sendVerificationCode(email string) {
	code := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	a.Redis.Client.Set(a.Redis.Ctx, "verificationCode:"+email, code, time.Minute*10)
	go a.sendEmail(email, code)
}
