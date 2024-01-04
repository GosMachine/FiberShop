package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (a *Handle) HandleHome(c *fiber.Ctx) error {
	return renderTemplate(c, "index", a, fiber.Map{"Title": "FiberShop"})
}

//func (a *Handle) HandleEmail(c *fiber.Ctx) error {
//	postCode := c.FormValue("code")
//	email := c.FormValue("email")
//	code := a.Redis.Client.Get(a.Redis.Ctx, "verificationCode:"+email).String()
//	if postCode != code {
//		return renderTemplate(c, "email", a, fiber.Map{"WrongCode": true, "Email": email})
//	}
//	switch c.OriginalURL() {
//	case "/register":
//		return a.authFinish(c, email)
//	case "/login":
//		return a.authFinish(c, email)
//	}
//	return nil
//}
