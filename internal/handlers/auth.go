package handlers

import (
	"FiberShop/web/view/auth"
	"context"
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
	email := a.Redis.GetToken(c.Cookies("token"))
	if email != "" {
		return c.Redirect("/")
	}
	return a.renderTemplate(c, auth.Login(a.getData(c, "Log In")))
}

func (a *Handle) HandleRegister(c *fiber.Ctx) error {
	email := a.Redis.GetToken(c.Cookies("token"))
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

func (a *Handle) HandleLogout(c *fiber.Ctx) error {
	go a.Grpc.Auth.Logout(context.Background(), c.Cookies("token"))
	c.Cookie(&fiber.Cookie{Name: "token", Secure: true, Value: "delete", Expires: time.Now().Add(-1 * time.Second)})
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
		token, err = a.Grpc.Auth.Register(context.Background(), data.Email, data.Password, c.IP(), data.Remember)
	case "login":
		token, err = a.Grpc.Auth.Login(context.Background(), data.Email, data.Password, c.IP(), data.Remember)
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
		a.Log.Error("auth error", zap.Error(err))
		return c.SendString("Internal error. Please try again.")
	}
	c.Cookie(&fiber.Cookie{Name: "token", Secure: true, Value: token, Expires: time.Now().Add(a.Redis.GetTokenTTL(token))})
	if action == "register" {
		a.sendVerificationCode(data.Email)
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
		return c.SendString("Internal error. Please try again.")
	}
	token, err := a.Grpc.Auth.OAuth(context.Background(), user.Email, c.IP())
	if err != nil {
		a.Log.Error("login error", zap.Error(err))
		return c.SendString("Internal error. Please try again.")
	}
	c.Cookie(&fiber.Cookie{Name: "token", Secure: true, Value: token, Expires: time.Now().Add(a.Redis.GetTokenTTL(token))})
	return c.Redirect("/")
}
