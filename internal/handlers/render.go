package handlers

import (
	"FiberShop/web/view/layout"
	"context"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"go.uber.org/zap"
)

func (a *Handle) getData(c *fiber.Ctx, title string) layout.Data {
	//todo наверное будет лучше передавать эмейл толлько там где он нужен
	email := a.Redis.GetToken(c.Cookies("token"))
	categories, err := a.Grpc.Product.GetCategoryNames(context.Background())
	if err != nil {
		a.Log.Error("error get categories name from grpc", zap.Error(err))
	}
	FinalData := layout.Data{
		Title:      title,
		Email:      email,
		Categories: categories,
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
