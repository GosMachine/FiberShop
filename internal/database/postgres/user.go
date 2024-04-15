package postgres

// import (
// 	"FiberShop/internal/database"
// 	"FiberShop/internal/models"
// 	"fmt"
// )

// // func (s *Storage) User(email string) (models.User, error) {
// // 	const op = "storage.postgres.User"

// // 	var user models.User
// // 	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
// // 		return models.User{}, fmt.Errorf("%s: %w", op, database.ErrUserNotFound)
// // 	}
// // 	return user, nil
// // }

// // func (s *Storage) UpdateUser(user models.User) error {
// // 	result := s.db.Save(&user)
// // 	if result.Error != nil {
// // 		return result.Error
// // 	}
// // 	return nil
// // }

// // func (s *Storage) UserPreload(tableName, email string, user *models.User) error {
// // 	err := s.db.Preload(tableName).Where("email = ?", email).First(&user).Error
// // 	if err != nil {
// // 		return err
// // 	}
// // 	return nil
// // }
