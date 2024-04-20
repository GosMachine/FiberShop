package handlers

import "github.com/gofiber/fiber/v2"

func (a *Handle) HandleCategory(c *fiber.Ctx) error {
	//todo продолжить тут, подключить сервис продуктов
	return c.Redirect("/")
}
