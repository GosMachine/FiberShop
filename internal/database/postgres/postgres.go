package postgres

import (
	"FiberShop/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New() (*Storage, error) {
	connection := "user=postgres password=postgres dbname=FiberShop host=127.0.0.1 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = database.AutoMigrate(models.User{}, models.Cart{}, models.Product{}, models.Category{}, models.Contact{}, models.CartItem{})
	if err != nil {
		return nil, err
	}
	return &Storage{db: database}, nil
}
