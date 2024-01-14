package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID     int64 `gorm:"primary_key"`
	UserID int64
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID"`
}

type CartItem struct {
	CartID    int64
	ProductID int64
	Quantity  int64
}
