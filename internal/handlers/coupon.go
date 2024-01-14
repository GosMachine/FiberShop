package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (a *Handle) HandleDiscountCoupon(c *fiber.Ctx) error {
	code := c.FormValue("code")
	coupon, err := a.Db.GetCoupon(code)
	if err != nil {
		a.Log.Error("err get coupon", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "InvalidCoupon"})
	}
	return c.JSON(fiber.Map{"type": coupon.ValueType, "value": coupon.Value})
}

func (a *Handle) HandleGiftCoupon(c *fiber.Ctx) error {
	return c.Redirect("/")
}
