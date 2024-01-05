package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (a *Handle) HandleHome(c *fiber.Ctx) error {
	return a.renderTemplate(c, "index", fiber.Map{"Title": "FiberShop"})
}

func (a *Handle) HandleAccountRecovery(c *fiber.Ctx) error {
	return a.renderTemplate(c, "account/recovery", fiber.Map{"Title": "Account recovery"})
}

func (a *Handle) HandleAccountRecoveryForm(c *fiber.Ctx) error {
	email := c.FormValue("email")
	_, err := a.Db.User(email)
	if err != nil {
		a.Log.Error("account_recovery error", zap.Error(err))
		return a.renderTemplate(c, "account/recovery", fiber.Map{"Title": "Account recovery", "Error": "UserIsNotFound"})
	}
	go func(email string) {
		a.sendEmail(email)
	}(email)
	return a.renderTemplate(c, "email", fiber.Map{"Title": "Email", "Email": email, "Action": "account_recovery"})
}
