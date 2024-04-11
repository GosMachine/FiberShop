package handlers

import (
	"FiberShop/internal/models"
	"FiberShop/internal/utils"
	"FiberShop/web/view/layout"
	"fmt"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"go.uber.org/zap"
)

func (a *Handle) getData(c *fiber.Ctx, title string) layout.Data {
	//todo мб не нужен баланс
	email, _ := utils.IsTokenValid(c.Cookies("token"))
	var (
		err  error
		user models.User
	)
	if email != "" {
		user, err = a.Redis.GetUserCache(email)
		if err != nil {
			a.Log.Error("error getting user", zap.Error(err))
		}
	}

	url := c.OriginalURL()
	ip := c.IP()
	a.Redis.IncrementViewCounter(url, ip+":"+url)
	viewers := a.Redis.Client.Get(a.Redis.Ctx, "viewers:"+url).Val()

	FinalData := layout.Data{
		Title:   title,
		Email:   email,
		Balance: fmt.Sprintf("%.2f", user.Balance),
		Viewers: viewers,
	}
	return FinalData
}

func (a *Handle) renderTemplate(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
