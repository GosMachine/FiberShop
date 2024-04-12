package postgres

import (
	"FiberShop/internal/models"
	"strconv"
)

func (s *Storage) DeleteItem(cart []models.CartItem, id string) ([]models.CartItem, error) {
	for i, item := range cart {
		if strconv.Itoa(int(item.ID)) == id {
			cart = append(cart[:i], cart[i+1:]...)
			break
		}
	}
	err := s.db.First("id = ?", id).Delete(models.CartItem{}).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}
