package services

import (
	"fidelizou-go/internal/database"
)

type Service struct {
	db database.Repository
}

func NewService(db database.Repository) Service {
	return Service{
		db: db,
	}
}

func (s *Service) Health() map[string]string {
	return s.db.Health()
}
