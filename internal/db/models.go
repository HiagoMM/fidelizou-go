// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type AuthProvider string

const (
	AuthProviderGOOLE AuthProvider = "GOOLE"
)

func (e *AuthProvider) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuthProvider(s)
	case string:
		*e = AuthProvider(s)
	default:
		return fmt.Errorf("unsupported scan type for AuthProvider: %T", src)
	}
	return nil
}

type NullAuthProvider struct {
	AuthProvider AuthProvider
	Valid        bool // Valid is true if AuthProvider is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuthProvider) Scan(value interface{}) error {
	if value == nil {
		ns.AuthProvider, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuthProvider.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuthProvider) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AuthProvider), nil
}

type Gender string

const (
	GenderMALE   Gender = "MALE"
	GenderFEMALE Gender = "FEMALE"
	GenderOTHER  Gender = "OTHER"
)

func (e *Gender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Gender(s)
	case string:
		*e = Gender(s)
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", src)
	}
	return nil
}

type NullGender struct {
	Gender Gender
	Valid  bool // Valid is true if Gender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullGender) Scan(value interface{}) error {
	if value == nil {
		ns.Gender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Gender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Gender), nil
}

type UserRole string

const (
	UserRolePROVIDER UserRole = "PROVIDER"
	UserRoleCLIENT   UserRole = "CLIENT"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole
	Valid    bool // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type FidelityCard struct {
	ID          int32
	Points      pgtype.Int4
	DatesOfUses pgtype.Timestamp
	UserID      pgtype.Int4
	ProgramID   pgtype.Int4
	CreatedAt   pgtype.Timestamp
}

type Program struct {
	ID                     int32
	Title                  pgtype.Text
	Description            pgtype.Text
	FidelityCardFrontImage pgtype.Text
	FidelityCardBackImage  pgtype.Text
	FidelityCardPointImage pgtype.Text
	FidelityCardMaxPoints  pgtype.Int4
	FinalDate              pgtype.Timestamp
	ActiveLinks            pgtype.Text
	OwnerID                pgtype.Int4
	CreatedAt              pgtype.Timestamp
}

type User struct {
	ID           int32
	Email        string
	Name         string
	Photo        string
	AuthProvider AuthProvider
	Credits      pgtype.Int4
	BirthDate    pgtype.Date
	Gender       NullGender
	UserRole     UserRole
	CreatedAt    pgtype.Timestamp
}