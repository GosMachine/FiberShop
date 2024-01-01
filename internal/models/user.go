package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID            int64  `gorm:"primary_key"`
	Email         string `gorm:"unique"`
	PassHash      []byte
	IpCreated     string
	LastLoginIp   string
	LastLoginDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	IsAdmin       bool
}
