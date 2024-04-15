package handlers

import (
	"FiberShop/internal/pkg/jwt"
	"FiberShop/web/view/layout"
	"fmt"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (a *Handle) getData(c *fiber.Ctx, title string) layout.Data {
	//todo в функциях где я получаю email, можно сначлао вызвыать getData и брать email оттуда
	timeStart := time.Now()

	email := jwt.IsTokenValid(c.Cookies("token"))

	fmt.Println(time.Since(timeStart))

	FinalData := layout.Data{
		Title: title,
		Email: email,
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

//todo разделить templ по функциям, а не держать вссе в одной
