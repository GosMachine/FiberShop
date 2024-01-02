package handlers

import "github.com/gofiber/fiber/v2"

func (a *Handle) HandleHome(c *fiber.Ctx) error {
	return renderTemplate(c, "index", a, fiber.Map{"Title": "FiberShop"})
}
