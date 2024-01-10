package handlers

import (
	"fmt"
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

func (a *Handle) HandleNotFound(c *fiber.Ctx) error {
	return a.renderTemplate(c, "404", fiber.Map{"Title": "Page not found"})
}

func (a *Handle) HandleContact(c *fiber.Ctx) error {
	return a.renderTemplate(c, "contact", fiber.Map{"Title": "Contact us"})
}

func (a *Handle) HandleContactForm(c *fiber.Ctx) error {
	type contactForm struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Message string `json:"message"`
	}
	var data contactForm
	if err := c.BodyParser(&data); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error create ticket",
		})
	}
	go func(data contactForm, ip string) {
		if err := a.Db.CreateTicket(data.Name, data.Email, data.Message, ip); err != nil {
			a.Log.Error("error create ticket", zap.Error(err))
		}
	}(data, c.IP())
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Ticket created",
	})
}
