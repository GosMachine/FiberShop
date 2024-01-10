package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          int64  `gorm:"primary_key"`
	Name        string `gorm:"unique"`
	Description string
	Price       float64
	Stock       int
	CategoryID  int64
	Category    Category
	ImageURL    string
}
