package handlers

import (
	"FiberShop/internal/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
)

func (a *Handle) HandleEmail(c *fiber.Ctx) error {
	postCode := c.FormValue("code")
	email := c.FormValue("email")
	action := c.FormValue("action")
	code, err := a.Redis.Client.Get(a.Redis.Ctx, "verificationCode:"+email).Result()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "CodeTimeError"})
	}
	if postCode != code {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "WrongCode"})
	}
	switch action {
	case "email_verification":
		go func(email string) {
			a.emailVerification(email)
		}(email)
	case "account_recovery":
		token, err := utils.NewToken(email, "", time.Hour*24)
		if err != nil {
			a.Log.Error("error create newToken", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalError"})
		}
		SetCookie("token", token, c, time.Now().Add(time.Hour*24))
		return a.renderTemplate(c, "account/change_pass", fiber.Map{"Title": "Change pass", "Email": email, "Action": action})
	case "change_email":
		newEmail, err := a.Redis.Client.Get(a.Redis.Ctx, "change_email:"+email).Result()
		if err != nil {
			a.Log.Error("error get newEmail", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalError"})
		}
		user, err := a.Redis.GetUserCache(email)
		if err != nil {
			a.Log.Error("error getting user cache", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalError"})
		}
		user.Email = newEmail
		err = a.Db.UpdateUser(user)
		if err != nil {
			a.Log.Error("error change email", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalError"})
		}

		if err := a.Redis.SetUserCache(user); err != nil {
			a.Redis.Log.Error("error set userCache", zap.Error(err))
		}
		a.Redis.Client.Del(a.Redis.Ctx, "change_email:"+email)
		token, err := utils.NewToken(newEmail, "on", time.Hour*336)
		if err != nil {
			a.Log.Error("error create newToken", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalError"})
		}
		SetCookie("token", token, c, time.Now().Add(time.Hour*336))
	}
	a.Redis.Client.Del(a.Redis.Ctx, "verificationCode:"+email)
	return c.Redirect("/account/settings")
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
	if err := a.Redis.SetUserCache(user); err != nil {
		a.Redis.Log.Error("error set userCache", zap.Error(err))
	}
}

func (a *Handle) HandleEmailResend(c *fiber.Ctx) error {
	email := c.FormValue("email")
	go func(email string) {
		a.sendEmail(email)
	}(email)
	return c.JSON(fiber.Map{"message": "Code sent successfully."})
}
