package postgres

import (
	"FiberShop/internal/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New() (*Storage, error) {
	connection := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	database, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = database.AutoMigrate(models.User{}, models.Product{},
		models.Category{}, models.Contact{}, models.CartItem{}, models.Coupon{})
	if err != nil {
		return nil, err
	}
	return &Storage{db: database}, nil
}
