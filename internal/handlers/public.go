package handlers

import "github.com/gofiber/fiber/v2"

func HandleHome(c *fiber.Ctx) error {
	return renderTemplate(c, "index")
}
