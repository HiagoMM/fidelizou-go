package models

import (
	"time"

	"github.com/google/uuid"
)

type Program struct {
	Id          uuid.UUID
	UserId      int
	Title       string
	Description string

	FidelityCardCoverImage string
	FidelityCardFrontImage string
	FidelityCardBackImage  string
	FidelityCardPointImage string
	FidelityCardMaxPoints  int

	FinalDate time.Time

	FidelityCards []FidelityCard
	ActiveLinks   []string
	CreatedAt     time.Time
}
