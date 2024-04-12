package postgres

import (
	"FiberShop/internal/models"
	"time"
)

func (s *Storage) GetCoupon(code string) (models.Coupon, error) {
	var coupon models.Coupon
	if err := s.db.Where("code = ?", code).First(&coupon).Error; err != nil {
		return models.Coupon{}, err
	}
	return coupon, nil
}

func (s *Storage) UpdateCoupon(coupon models.Coupon) error {
	err := s.db.Save(&coupon).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) CreateCoupon(code string, value float64, valueType models.ValueTypeEnum, maxUsageCount, usageCount int, expiresAt *time.Time) error {
	coupon := models.Coupon{
		Code:          code,
		Value:         value,
		ValueType:     valueType,
		MaxUsageCount: maxUsageCount,
		UsageCount:    usageCount,
		ExpiresAt:     expiresAt,
	}
	err := s.db.Create(&coupon).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteCoupon(code string) error {
	err := s.db.Where("code = ?", code).Delete(&models.Coupon{}).Error
	if err != nil {
		return err
	}
	return nil
}
