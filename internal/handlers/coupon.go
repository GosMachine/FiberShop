package handlers

//TODO перенос в микросервис корзины

import (
	"github.com/gofiber/fiber/v2"
)

func (a *Handle) HandleDiscountCoupon(c *fiber.Ctx) error {
	//TODO перенос в микросервис корзины
	// alert := alerts.Alert{
	// 	Name:    "Coupon",
	// 	Message: "Coupon not found",
	// 	Button:  "couponBtn",
	// }
	// code := c.FormValue("code")
	// coupon, err := a.Db.GetCoupon(code)
	// if err != nil || coupon.Type != models.Discount {
	// 	a.Log.Error("err get coupon", zap.Error(err))
	// 	return a.renderTemplate(c, alerts.Error(alert, a.getData(c, "Your Cart")))
	// }
	// if coupon.UsageCount >= coupon.MaxUsageCount || coupon.ExpiresAt.Before(time.Now()) {
	// 	alert.Message = "Coupon has expired"
	// 	return a.renderTemplate(c, alerts.Error(alert, a.getData(c, "Your Cart")))
	// }
	// coupon.UsageCount++
	// if err = a.Db.UpdateCoupon(coupon); err != nil {
	// 	alert.Name = "Internal server error"
	// 	alert.Message = "Please try again."
	// 	a.Log.Error("error update coupon usage count", zap.Error(err))
	// 	return a.renderTemplate(c, alerts.Error(alert, a.getData(c, "Your Cart")))
	// }
	// alert.Message = "Successfully activated"
	// c.Set("HX-Retarget", "#couponResponse")
	// return c.SendString(fmt.Sprintf("<script>percentage = %d;updateTotalCartPrice()</script>", 10))
	return c.Redirect("/")
}

func (a *Handle) HandleGiftCoupon(c *fiber.Ctx) error {
	return c.Redirect("/")
}

// func applyCoupon(totalPrice float64, coupon *models.Coupon) float64 {
// 	switch coupon.ValueType {
// 	case models.FixedAmount:
// 		return totalPrice - coupon.Value
// 	case models.Percentage:
// 		return totalPrice * (1 - coupon.Value)
// 	default:
// 		return totalPrice
// 	}
// }
