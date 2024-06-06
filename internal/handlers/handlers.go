package handlers

import (
	"fidelizou-go/internal/db"
	"fidelizou-go/internal/services"
)

type Handlers struct {
	s services.Service
}

func NewHandlers(repo *db.Queries) *Handlers {
	return &Handlers{
		s: services.NewService(repo),
	}
}
