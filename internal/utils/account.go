package utils

import "FiberShop/internal/models"

func applyCoupon(totalPrice float64, coupon *models.Coupon) float64 {
	switch coupon.ValueType {
	case models.FixedAmount:
		return totalPrice - coupon.Value
	case models.Percentage:
		return totalPrice * (1 - coupon.Value)
	default:
		return totalPrice
	}
}
