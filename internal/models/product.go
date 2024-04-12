package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          int    `gorm:"primary_key"`
	Name        string `gorm:"unique"`
	Description string
	Price       float64
	Stock       int
	CategoryID  int
	Category    Category
	ImageURL    string
}
