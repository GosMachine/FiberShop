package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	ID        int
	UserID    int
	Product   Product `gorm:"foreignKey:ProductID"`
	ProductID int
	Quantity  int
}
