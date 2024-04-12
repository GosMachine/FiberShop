package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	ID      int    `gorm:"primary_key"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	IP      string `json:"ip"`
}
