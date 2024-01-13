package handlers

import (
	"FiberShop/internal/utils"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type RequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember string `json:"remember"`
}

func (a *Handle) HandleLogin(c *fiber.Ctx) error {
	email, _ := utils.IsTokenValid(c.Cookies("token"))
	if email != "" {
		return c.Redirect("/")
	}
	return a.renderTemplate(c, "account/login", fiber.Map{"Title": "Log In"})
}

func (a *Handle) HandleRegister(c *fiber.Ctx) error {
	email, _ := utils.IsTokenValid(c.Cookies("token"))
	if email != "" {
		return c.Redirect("/")
	}
	return a.renderTemplate(c, "account/register", fiber.Map{"Title": "Sign Up"})
}

func (a *Handle) HandleLoginForm(c *fiber.Ctx) error {
	return a.auth(c, "login")
}

func (a *Handle) HandleRegisterForm(c *fiber.Ctx) error {
	return a.auth(c, "register")
}

func HandleLogout(c *fiber.Ctx) error {
	SetCookie("token", "delete", c, time.Now().Add(-1*time.Second))
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
		token, err = a.Client.Login(context.Background(), data.Email, data.Password, c.IP(), data.Remember, "default")
	}
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.InvalidArgument {
				a.Log.Error(action+" error", zap.Error(err))
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "InvalidCredentials"})
			} else if st.Code() == codes.AlreadyExists {
				a.Log.Error(action+" error", zap.Error(err))
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "AlreadyExists"})
			}
		}
		a.Log.Error("login error", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "InternalError"})
	}
	expires := time.Now().Add(time.Hour * 24)
	if data.Remember == "on" {
		expires = time.Now().Add(time.Hour * 336)
	}
	SetCookie("token", token, c, expires)
	if action == "register" {
		go func(data RequestData) {
			a.sendEmail(data.Email)
		}(data)
		return a.renderTemplate(c, "email", fiber.Map{"Title": "Email", "Email": data.Email, "Action": "email_verification"})
	}
	return c.Redirect("/")
}

func (a *Handle) HandleOAuthCallback(ctx *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(ctx)
	if err != nil {
		return err
	}
	token, err := a.Client.Login(context.Background(), user.Email, "", ctx.IP(), "on", "OAuth")
	if err != nil {
		a.Log.Error("login error", zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalError"})
	}
	expires := time.Now().Add(time.Hour * 336)

	SetCookie("token", token, ctx, expires)
	return ctx.Redirect("/")
}
