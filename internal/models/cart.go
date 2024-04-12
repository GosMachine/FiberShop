package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	UserID    int
	Product   Product `gorm:"foreignKey:ProductID"`
	ProductID int
	Quantity  int
}
