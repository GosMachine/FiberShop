package handlers

import (
	"FiberShop/web/view/layout"
	"context"

	"github.com/gofiber/fiber/v2"
)

func (a *Handle) HandleCategory(c *fiber.Ctx) error {
	categoryName := c.Params("category")
	_, err := a.Grpc.Product.GetCategory(context.Background(), categoryName)
	if err != nil {
		return a.renderTemplate(c, layout.NotFound(a.getData(c, categoryName)))
	}
	//todo тут нужно сделать темплейт и его рендерить (страница категории)
	return c.Redirect("/")
}
