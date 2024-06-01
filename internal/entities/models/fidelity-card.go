package models

import (
	"time"

	"github.com/google/uuid"
)

type FidelityCard struct {
	Id        uuid.UUID
	Points    int
	UserId    int
	ProgramId int

	User    User
	Program Program

	DatesOfUses []time.Time
	CreatedAt   time.Time
}
