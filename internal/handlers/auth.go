package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type RequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember string `json:"remember"`
}

func HandleLogin(c *fiber.Ctx) error {
	return renderTemplate(c, "account/login")
}

func HandleRegister(c *fiber.Ctx) error {
	return renderTemplate(c, "account/register")
}

func (a *Handle) HandleLoginForm(c *fiber.Ctx) error {
	return a.auth(c, "login")
}

func (a *Handle) HandleRegisterForm(c *fiber.Ctx) error {
	return a.auth(c, "register")
}

func (a *Handle) auth(c *fiber.Ctx, action string) error {
	var (
		token string
		err   error
		data  RequestData
	)
	if err := c.BodyParser(&data); err != nil {
		a.Log.Error("bodyParse error", zap.Error(err))
		return err
	}
	switch action {
	case "register":
		token, err = a.Client.Register(context.Background(), data.Email, data.Password, c.IP(), data.Remember)
	case "login":
		token, err = a.Client.Login(context.Background(), data.Email, data.Password, c.IP(), data.Remember)
	}
	if err != nil {
		a.Log.Error(action+" error", zap.Error(err))
		return err
	}
	cookie := fiber.Cookie{
		Name:   "token",
		Secure: true,
		Value:  token,
	}
	c.Cookie(&cookie)
	return c.Redirect("/")
}
