package models

import "fidelizou-go/internal/db"

type Session struct {
	Id   int32
	Role db.UserRole
}
