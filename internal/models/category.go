package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID          int64  `gorm:"primary_key"`
	Name        string `gorm:"unique"`
	Description string
	Products    []Product
}
