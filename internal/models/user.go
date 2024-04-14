package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID            int    `gorm:"primary_key"`
	Email         string `gorm:"unique_index"`
	EmailVerified bool   `gorm:"index"`
	PassHash      []byte
	IpCreated     string
	LastLoginIp   string
	LastLoginDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
