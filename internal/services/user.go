package services

import (
	"fidelizou-go/internal/entities/err"
	"fidelizou-go/internal/entities/models"
)

var (
	ErrInvalidRole = err.NewBadRequest("Invalid role")
)

func (s *Service) GetUser(email string) (models.User, error) {
	return s.db.GetUser(email)
}

func (s *Service) ToggleUserRole(session *models.Session) error {

	var newRole int8

	switch session.Role {
	case models.RoleClient:
		newRole = models.RoleProvider
	case models.RoleProvider:
		newRole = models.RoleClient
	default:
		return ErrInvalidRole
	}

	if err := s.db.UpdateUserRole(newRole, session.Email); err != nil {
		return err
	}
	return nil
}
