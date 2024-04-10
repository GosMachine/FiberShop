package handlers

import (
	"FiberShop/internal/utils"
	"FiberShop/web/view/alerts"
	"FiberShop/web/view/auth"
	"FiberShop/web/view/contact"
	"FiberShop/web/view/index"
	"FiberShop/web/view/layout"
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (a *Handle) HandleHome(c *fiber.Ctx) error {
	return a.renderTemplate(c, index.Show(a.getData(c, "GosBoost")))
}

func (a *Handle) HandleAccountRecovery(c *fiber.Ctx) error {
	email, _ := utils.IsTokenValid(c.Cookies("token"))
	if email != "" {
		return c.Redirect("/")
	}
	return a.renderTemplate(c, auth.Recovery(a.getData(c, "Account recovery")))
}

func (a *Handle) HandleAccountRecoveryForm(c *fiber.Ctx) error {
	var data RequestData
	if err := c.BodyParser(&data); err != nil {
		a.Log.Error("bodyParse error", zap.Error(err))
		return c.SendString("Internal error. Please try again.")
	}
	_, err := a.Db.User(data.Email)
	if err != nil {
		a.Log.Error("account_recovery error", zap.Error(err))
		return c.SendString("User is not found")
	}
	code := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	go func(email string) {
		a.sendEmail(email, code)
	}(data.Email)
	c.Set("HX-Redirect", "/email?action=account_recovery&address="+data.Email)
	return c.SendStatus(200)
}

func (a *Handle) HandleNotFound(c *fiber.Ctx) error {
	return a.renderTemplate(c, layout.NotFound(a.getData(c, "Page not found")))
}

func (a *Handle) HandleContact(c *fiber.Ctx) error {
	return a.renderTemplate(c, contact.Show(a.getData(c, "Contact us")))
}

func (a *Handle) HandleContactForm(c *fiber.Ctx) error {
	type contactForm struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Message string `json:"message"`
	}
	var data contactForm
	if err := c.BodyParser(&data); err != nil {
		a.Log.Error("error create ticket", zap.Error(err))
		return a.renderTemplate(c, alerts.Error("Ticket", a.getData(c, "Contact us")))
	}
	if err := a.Db.CreateTicket(data.Name, data.Email, data.Message, c.IP()); err != nil {
		a.Log.Error("error create ticket", zap.Error(err))
		return a.renderTemplate(c, alerts.Error("Ticket", a.getData(c, "Contact us")))
	}
	return a.renderTemplate(c, alerts.Success("Ticket", a.getData(c, "Contact us")))
}
