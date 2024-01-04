package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
)

type RequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember string `json:"remember"`
}

func (a *Handle) HandleLogin(c *fiber.Ctx) error {
	return renderTemplate(c, "account/login", a, fiber.Map{"Title": "Log In"})
}

func (a *Handle) HandleRegister(c *fiber.Ctx) error {
	return renderTemplate(c, "account/register", a, fiber.Map{"Title": "Sign Up"})
}

func (a *Handle) HandleLoginForm(c *fiber.Ctx) error {
	return a.auth(c, "login")
}

func (a *Handle) HandleRegisterForm(c *fiber.Ctx) error {
	return a.auth(c, "register")
}

func HandleLogout(c *fiber.Ctx) error {
	setCookie("token", "delete", c, time.Now().Add(-1*time.Second))
	return c.Redirect("/")
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
	expires := time.Now().Add(time.Hour * 24)
	if data.Remember == "on" {
		expires = time.Now().Add(time.Hour * 336)
	}
	setCookie("token", token, c, expires)
	if action == "register" {
		go func(data RequestData) {
			a.sendEmail(data.Email)
		}(data)
		return renderTemplate(c, "email", a, fiber.Map{"Email": data.Email, "Action": "register"})
	}
	return c.Redirect("/")
}
