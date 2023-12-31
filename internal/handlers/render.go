package handlers

import "github.com/gofiber/fiber/v2"

func renderTemplate(c *fiber.Ctx, tmpl string, layouts ...string) error {
	return c.Render(tmpl, fiber.Map{}, layouts...)
}
