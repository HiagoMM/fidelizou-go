package services

import "fidelizou-go/internal/db"

type Service struct {
	db *db.Queries
}

func NewService(db *db.Queries) Service {
	return Service{
		db: db,
	}
}
