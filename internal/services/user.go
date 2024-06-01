package services

import "fidelizou-go/internal/entities/models"

func (s *Service) GetUser(email string) (models.User, error) {
	return s.db.GetUser(email)
}

func (s *Service) ToggleUserRole(session *models.Session) error {
	err := s.db.ToggleUserRole(session.Email)
	if err != nil {
		return err
	}
	return nil
}
