package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID            int64  `gorm:"primary_key"`
	Email         string `gorm:"unique"`
	EmailVerified bool
	Balance       float64
	PassHash      []byte
	IpCreated     string
	LastLoginIp   string
	LastLoginDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CartID        int64
	Cart          Cart `json:"cart" gorm:"foreignKey:CartID"`
}

func (User) Indexes() []string {
	return []string{"balance"}
}
