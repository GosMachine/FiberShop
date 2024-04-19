package postgres

// import (
// 	"FiberShop/internal/models"
// )

// func (s *Storage) CreateCategory(name, description string) (models.Category, error) {
// 	category := models.Category{Name: name, Description: description, Products: []models.Product{}}

// 	if err := s.db.Create(&category).Error; err != nil {
// 		return models.Category{}, err
// 	}
// 	return category, nil
// }
