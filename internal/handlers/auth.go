package handlers

import "github.com/gofiber/fiber/v2"

func HandleLogin(c *fiber.Ctx) error {
	return renderTemplate(c, "account/login")
}
func HandleRegister(c *fiber.Ctx) error {
	return renderTemplate(c, "account/register")
}
