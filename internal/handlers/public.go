package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (a *Handle) HandleHome(c *fiber.Ctx) error {
	return renderTemplate(c, "index", a, fiber.Map{"Title": "FiberShop"})
}

func (a *Handle) HandleAccountRecovery(c *fiber.Ctx) error {
	return renderTemplate(c, "account/recovery", a, fiber.Map{"Title": "Account recovery"})
}
