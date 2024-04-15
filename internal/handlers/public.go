package handlers

import (
	"FiberShop/web/view/auth"
	"FiberShop/web/view/index"
	"FiberShop/web/view/layout"
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (a *Handle) HandleHome(c *fiber.Ctx) error {
	return a.renderTemplate(c, index.Show(a.getData(c, "GosBoost")))
}

func (a *Handle) HandleAccountRecovery(c *fiber.Ctx) error {
	data := a.getData(c, "Account recovery")
	if data.Email != "" {
		return c.Redirect("/")
	}
	return a.renderTemplate(c, auth.Recovery(data))
}

func (a *Handle) HandleAccountRecoveryForm(c *fiber.Ctx) error {
	var data RequestData
	if err := c.BodyParser(&data); err != nil {
		a.Log.Error("bodyParse error", zap.Error(err))
		return c.SendString("Internal error. Please try again.")
	}
	_, err := a.Client.EmailVerified(context.Background(), data.Email)
	if err != nil {
		a.Log.Error("account_recovery error", zap.Error(err))
		return c.SendString("User is not found")
	}
	code := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	a.Redis.Client.Set(a.Redis.Ctx, "verificationCode:"+data.Email, code, time.Minute*10)
	go func(email string) {
		a.sendEmail(email, code)
	}(data.Email)
	c.Set("HX-Redirect", "/email?action=account_recovery&address="+data.Email)
	return c.SendStatus(200)
}

func (a *Handle) HandleNotFound(c *fiber.Ctx) error {
	return a.renderTemplate(c, layout.NotFound(a.getData(c, "Page not found")))
}
