package postgres

import (
	"FiberShop/internal/models"
	"fmt"
)

func (s *Storage) GetTicket(id int64) (models.Contact, error) {
	var contact models.Contact
	if err := s.db.Where("id = ?", id).First(&contact).Error; err != nil {
		return models.Contact{}, err
	}
	return contact, nil
}

func (s *Storage) CreateTicket(name, email, message, ip string) error {
	const op = "storage.postgres.CreateTicket"
	contact := models.Contact{Name: name, Email: email, Message: message, IP: ip}
	if err := s.db.Create(&contact).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
