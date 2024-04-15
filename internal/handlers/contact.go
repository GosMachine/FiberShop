package handlers

import (
	"FiberShop/web/view/alerts"
	"FiberShop/web/view/contact"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type contactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (a *Handle) HandleContact(c *fiber.Ctx) error {
	return a.renderTemplate(c, contact.Show(a.getData(c, "Contact us")))
}

func (a *Handle) HandleContactForm(c *fiber.Ctx) error {
	alert := alerts.Alert{
		Name:    "Ticket",
		Message: "Error create ticket",
		Button:  "submitBtn",
	}
	var data contactForm
	if err := c.BodyParser(&data); err != nil {
		a.Log.Error("error create ticket", zap.Error(err))
		return a.renderTemplate(c, alerts.Error(alert, a.getData(c, "Contact us")))
	}
	if err := a.Db.CreateTicket(data.Name, data.Email, data.Message, c.IP()); err != nil {
		a.Log.Error("error create ticket", zap.Error(err))
		return a.renderTemplate(c, alerts.Error(alert, a.getData(c, "Contact us")))
	}
	alert.Message = "Successfully created"
	return a.renderTemplate(c, alerts.Success(alert, a.getData(c, "Contact us")))
}
