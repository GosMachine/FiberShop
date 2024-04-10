package handlers

import (
	"FiberShop/internal/utils"
	"FiberShop/web/view/auth"
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RequestData struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Remember        string `json:"remember"`
}

func (a *Handle) HandleLogin(c *fiber.Ctx) error {
	email, _ := utils.IsTokenValid(c.Cookies("token"))
	if email != "" {
		return c.Redirect("/")
	}
	return a.renderTemplate(c, auth.Login(a.getData(c, "Log In")))
}

func (a *Handle) HandleRegister(c *fiber.Ctx) error {
	email, _ := utils.IsTokenValid(c.Cookies("token"))
	if email != "" {
		return c.Redirect("/")
	}
	return a.renderTemplate(c, auth.Register(a.getData(c, "Sign Up")))
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
		return c.SendString("Internal error. Please try again.")
	}
	switch action {
	case "register":
		if data.Password != data.ConfirmPassword {
			return c.SendString("Password mismatch.")
		}
		token, err = a.Client.Register(context.Background(), data.Email, data.Password, c.IP(), data.Remember)
	case "login":
		token, err = a.Client.Login(context.Background(), data.Email, data.Password, c.IP(), data.Remember, "default")
	}
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.InvalidArgument {
				a.Log.Error(action+" error", zap.Error(err))
				return c.SendString("Invalid email or password.")
			} else if st.Code() == codes.AlreadyExists {
				a.Log.Error(action+" error", zap.Error(err))
				return c.SendString("User already exists.")
			}
		}
		a.Log.Error("login error", zap.Error(err))
		return c.SendString("Internal error. Please try again.")
	}
	expires := time.Now().Add(time.Hour * 24)
	if data.Remember == "on" {
		expires = time.Now().Add(time.Hour * 336)
	}
	SetCookie("token", token, c, expires)
	if action == "register" {
		code := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
		go func(data RequestData) {
			a.sendEmail(data.Email, code)
		}(data)
		c.Set("HX-Redirect", "/email?action=email_verification&address="+data.Email)
		return c.SendStatus(200)
	}
	c.Set("HX-Redirect", "/")
	return c.SendStatus(200)
}

func (a *Handle) HandleOAuthCallback(c *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		a.Log.Error("login error", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Internal error. Please try again.")
	}
	token, err := a.Client.Login(context.Background(), user.Email, "", c.IP(), "on", "OAuth")
	if err != nil {
		a.Log.Error("login error", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Internal error. Please try again.")
	}
	expires := time.Now().Add(time.Hour * 336)

	SetCookie("token", token, c, expires)
	return c.Redirect("/")
}
