package handlers

import (
	"FiberShop/internal/utils"
	"FiberShop/web/view/account"
	"FiberShop/web/view/alerts"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (a *Handle) HandleAccountCart(c *fiber.Ctx) error {
	email, _ := utils.IsTokenValid(c.Cookies("token"))
	user, err := a.Redis.GetUserCache(email)
	if err != nil {
		a.Log.Error("error get user cache", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("internal error")
	}
	var totalCartPrice float64
	for _, item := range user.Cart {
		totalCartPrice += item.Product.Price * float64(item.Quantity)
	}
	return a.renderTemplate(c, account.Cart(a.getData(c, "Your Cart"), user.Cart, totalCartPrice))
}

func (a *Handle) HandleDeleteItem(c *fiber.Ctx) error {
	email, _ := utils.IsTokenValid(c.Cookies("token"))
	id := c.Query("id")
	user, err := a.Redis.GetUserCache(email)
	if err != nil {
		a.Log.Error("error get user cache", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("internal error")
	}
	user.Cart, err = a.Db.DeleteItem(user.Cart, id)
	if err != nil {
		a.Log.Error("error delete item", zap.Error(err))
		alert := alerts.Alert{
			Name:    "Error",
			Message: "Deleting error",
			Button:  "remove" + id,
		}
		return a.renderTemplate(c, alerts.Error(alert, a.getData(c, "Error")))
	}
	if err := a.Redis.SetUserCache(user); err != nil {
		a.Log.Error("error set user cache", zap.Error(err))
	}
	return c.SendStatus(200)
}
