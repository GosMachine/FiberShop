package handlers

import (
	"FiberShop/internal/utils"
	"FiberShop/web/view/layout"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (a *Handle) getData(c *fiber.Ctx, title string) layout.Data {
	//todo мб не нужен баланс
	email, _ := utils.IsTokenValid(c.Cookies("token"))
	url := c.OriginalURL()
	ip := c.IP()
	a.Redis.IncrementViewCounter(url, ip+":"+url)
	viewers := a.Redis.Client.Get(a.Redis.Ctx, "viewers:"+url).Val()

	FinalData := layout.Data{
		Title:   title,
		Email:   email,
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
