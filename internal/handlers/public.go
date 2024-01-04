package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (a *Handle) HandleHome(c *fiber.Ctx) error {
	return a.renderTemplate(c, "index", fiber.Map{"Title": "FiberShop"})
}

func (a *Handle) HandleAccountRecovery(c *fiber.Ctx) error {
	return a.renderTemplate(c, "account/recovery", fiber.Map{"Title": "Account recovery"})
}
