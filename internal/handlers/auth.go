package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

type RequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember string `json:"remember"`
}

func (a *Handle) HandleLogin(c *fiber.Ctx) error {
	return a.renderTemplate(c, "account/login", fiber.Map{"Title": "Log In"})
}

func (a *Handle) HandleRegister(c *fiber.Ctx) error {
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
				return a.renderTemplate(c, "account/"+action, fiber.Map{"Title": strings.Title(action), "Error": "InvalidCredentials"})
			} else if st.Code() == codes.AlreadyExists {
				a.Log.Error(action+" error", zap.Error(err))
				return a.renderTemplate(c, "account/register", fiber.Map{"Title": "Register", "Error": "AlreadyExists"})
			}
		}
		a.Log.Error("login error", zap.Error(err))
		return a.renderTemplate(c, "account/"+action, fiber.Map{"Title": strings.Title(action), "Error": "InternalError"})
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
		return a.renderTemplate(c, "email", fiber.Map{"Title": "Email", "Email": data.Email, "Action": "register"})
	}
	return c.Redirect("/")
}

func (a *Handle) HandleChangePassForm(c *fiber.Ctx) error {
	pass := c.FormValue("password")
	email := c.FormValue("email")
	action := c.FormValue("action")
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		a.Log.Error("failed to generate password hash", zap.Error(err))
	}
	user, err := a.Redis.GetUserCache(email)
	if err != nil {
		a.Log.Error("failed to get user cache", zap.Error(err))
	}
	user.PassHash = passHash
	user.LastLoginIp = c.IP()
	user.LastLoginDate = time.Now()
	err = a.Db.UpdateUser(user)
	if err != nil {
		a.Log.Error("failed to update user", zap.Error(err))
	}
	a.Log.Info("password changed successfully", zap.String("email", email))
	if action == "change_pass" {
		return c.Redirect("/account")
	}
	return c.Redirect("/")
}

func (a *Handle) HandleAuthCallback(ctx *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(ctx)
	if err != nil {
		return err
	}
	token, err := a.Client.Login(context.Background(), user.Email, "", ctx.IP(), "on", "OAuth")
	if err != nil {
		a.Log.Error("login error", zap.Error(err))
		return a.renderTemplate(ctx, "account/login", fiber.Map{"Title": "Log In", "Error": "InternalError"})
	}
	expires := time.Now().Add(time.Hour * 336)

	SetCookie("token", token, ctx, expires)
	return ctx.Redirect("/")
}
