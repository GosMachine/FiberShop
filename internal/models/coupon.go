package models

import (
	"gorm.io/gorm"
	"time"
)

type ValueTypeEnum string

type CouponType string

const (
	FixedAmount ValueTypeEnum = "fixed_amount"
	Percentage  ValueTypeEnum = "percentage"
	Discount    CouponType    = "discount"
	Gift        CouponType    = "gift"
)

type Coupon struct {
	gorm.Model
	Type          CouponType
	Code          string        `gorm:"uniqueIndex;not null"`
	ValueType     ValueTypeEnum `gorm:"not null"`
	Value         float64       `gorm:"not null"`
	MaxUsageCount int
	UsageCount    int
	ExpiresAt     *time.Time
}
