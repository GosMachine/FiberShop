package postgres

import (
	"FiberShop/internal/database"
	"FiberShop/internal/models"
	"fmt"
)

func (s *Storage) User(email string) (models.User, error) {
	const op = "storage.postgres.User"

	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, database.ErrUserNotFound)
	}
	return user, nil
}
