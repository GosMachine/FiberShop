package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (a *Handle) HandleEmail(c *fiber.Ctx) error {
	postCode := c.FormValue("code")
	email := c.FormValue("email")
	action := c.FormValue("action")
	code := a.Redis.Client.Get(a.Redis.Ctx, "verificationCode:"+email).Val()
	if postCode != code {
		return a.renderTemplate(c, "email", fiber.Map{"WrongCode": true, "Email": email, "Action": action})
	}
	switch action {
	case "register":
		go func(email string) {
			a.emailVerification(email)
		}(email)
	case "email_verification":
		go func(email string) {
			a.emailVerification(email)
		}(email)
	case "account_recovery":
	case "change_password":
	}

	return c.Redirect("/")
}

func (a *Handle) emailVerification(email string) {
	user, err := a.Redis.GetUserCache(email)
	if err != nil {
		return
	}
	user.EmailVerified = true
	err = a.Db.UpdateUser(user)
	if err != nil {
		a.Log.Error("error email verification", zap.Error(err))
	}
}
