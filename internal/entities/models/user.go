package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	GenderMale   = 1
	GenderFemale = 2
	GenderOther  = 3
)
const (
	RoleProvider = 1
	RoleClient   = 2
)
const (
	SignInMethodGoogle   = 1
	SignInMethodFacebook = 2
)

type User struct {
	Id    uuid.UUID
	Email string
	Name  string
	Photo string

	Credits    int
	TopCredits int

	SignInMethod uint8
	Gender       uint8
	Role         uint8

	Programs      []Program
	FidelityCards []FidelityCard

	BirthDate time.Time
	CreatedAt time.Time
}

type Session struct {
	Email string
	Role  uint8
}
