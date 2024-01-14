package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    int64
	CartItems []CartItem
}

type CartItem struct {
	gorm.Model
	CartID     uint
	Product    Product `gorm:"foreignKey:ProductID"`
	ProductID  int64
	Quantity   int64
	TotalPrice float64
}
