package services

import (
	"context"
	"fidelizou-go/internal/db"
	"fidelizou-go/internal/entities/err"
	"fidelizou-go/internal/entities/models"
)

var (
	ErrInvalidRole = err.NewBadRequest("Invalid role")
)

func (s *Service) GetUser(ctx context.Context, email string) (db.User, error) {
	return s.db.GetUserByEmail(ctx, email)
}

func (s *Service) ToggleUserRole(ctx context.Context, session *models.Session) error {

	var newRole db.UserRole

	switch session.Role {
	case db.UserRoleCLIENT:
		newRole = db.UserRolePROVIDER
	case db.UserRolePROVIDER:
		newRole = db.UserRoleCLIENT
	default:
		return ErrInvalidRole
	}

	updateParams := db.UpdateUserRoleParams{
		UserRole: newRole,
		ID:       session.Id,
	}

	if err := s.db.UpdateUserRole(ctx, updateParams); err != nil {
		return err
	}
	return nil
}
