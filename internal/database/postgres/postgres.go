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
	connection := "user=FiberShop password=FiberShop dbname=FiberShop host=127.0.0.1 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = database.AutoMigrate(models.User{}, models.Product{}, models.Category{})
	if err != nil {
		return nil, err
	}
	return &Storage{db: database}, nil
}
